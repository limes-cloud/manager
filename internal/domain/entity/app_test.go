package entity

import (
	"testing"

	"github.com/limes-cloud/kratosx/pkg/value"
)

func TestAppParseSetting(t *testing.T) {
	tests := []struct {
		name        string
		setting     string
		wantErr     bool
		wantModules []string
		wantSecret  string
	}{
		{
			name:        "空配置返回零值且无错误",
			setting:     "",
			wantModules: nil,
		},
		{
			name:        "完整配置正常解析",
			setting:     `{"jwt":{"secret":"s1","expire":86400},"modules":["configure","cron"]}`,
			wantModules: []string{"configure", "cron"},
			wantSecret:  "s1",
		},
		{
			name:    "非法JSON返回错误",
			setting: `{"jwt":`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{Keyword: "manager", Setting: tt.setting}
			setting, err := app.ParseSetting()
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseSetting() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(setting.Modules) != len(tt.wantModules) {
				t.Fatalf("Modules = %v, want %v", setting.Modules, tt.wantModules)
			}
			for i, m := range tt.wantModules {
				if setting.Modules[i] != m {
					t.Errorf("Modules[%d] = %s, want %s", i, setting.Modules[i], m)
				}
			}
			if setting.JWT.Secret != tt.wantSecret {
				t.Errorf("JWT.Secret = %s, want %s", setting.JWT.Secret, tt.wantSecret)
			}
		})
	}
}

func TestAppScopeModules(t *testing.T) {
	tests := []struct {
		name    string
		keyword string
		setting string
		want    []string
	}{
		{
			name:    "无订阅时仅包含自身",
			keyword: "manager",
			setting: `{}`,
			want:    []string{"manager"},
		},
		{
			name:    "自身加订阅模块",
			keyword: "manager",
			setting: `{"modules":["configure","cron"]}`,
			want:    []string{"manager", "configure", "cron"},
		},
		{
			name:    "订阅列表包含自身时去重",
			keyword: "manager",
			setting: `{"modules":["manager","configure"]}`,
			want:    []string{"manager", "configure"},
		},
		{
			name:    "过滤空字符串",
			keyword: "manager",
			setting: `{"modules":["","configure"]}`,
			want:    []string{"manager", "configure"},
		},
		{
			name:    "配置解析失败时退化为仅自身",
			keyword: "manager",
			setting: `{invalid`,
			want:    []string{"manager"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{Keyword: tt.keyword, Setting: tt.setting}
			got := app.ScopeModules()
			if len(got) != len(tt.want) {
				t.Fatalf("ScopeModules() = %v, want %v", got, tt.want)
			}
			for i, w := range tt.want {
				if got[i] != w {
					t.Errorf("ScopeModules()[%d] = %s, want %s", i, got[i], w)
				}
			}
		})
	}
}

func TestAppHasModule(t *testing.T) {
	app := &App{Keyword: "manager", Setting: `{"modules":["configure"]}`}

	if !app.HasModule("manager") {
		t.Error("HasModule(自身) = false, want true")
	}
	if !app.HasModule("configure") {
		t.Error("HasModule(已订阅) = false, want true")
	}
	if app.HasModule("cron") {
		t.Error("HasModule(未订阅) = true, want false")
	}
	if app.HasModule("") {
		t.Error("HasModule(空) = true, want false")
	}
}

func TestAppIsLoginable(t *testing.T) {
	if (&App{}).IsLoginable() {
		t.Error("Loginable 为 nil 时应视为不可登录")
	}
	if (&App{Loginable: value.Pointer(false)}).IsLoginable() {
		t.Error("Loginable = false 时应不可登录")
	}
	if !(&App{Loginable: value.Pointer(true)}).IsLoginable() {
		t.Error("Loginable = true 时应可登录")
	}
}

func TestAppIsBase(t *testing.T) {
	if !(&App{Type: AppTypeBase}).IsBase() {
		t.Error("type=base 应为基础设施")
	}
	if (&App{Type: AppTypeApp}).IsBase() {
		t.Error("type=app 不应为基础设施")
	}
}
