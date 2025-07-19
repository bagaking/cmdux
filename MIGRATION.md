# ccmodel 到 cmdux 迁移指南

## 概述

本指南演示如何将 ccmodel 项目中的 UI/UX 代码迁移到独立的 cmdux 库。

## 迁移对比

### 原 ccmodel 代码 (internal/ui/ui.go)

```go
// 原代码
import "github.com/fatih/color"

var (
    Primary   = color.New(color.FgHiCyan, color.Bold)
    Secondary = color.New(color.FgHiBlue)
    Success   = color.New(color.FgHiGreen, color.Bold)
)

func Header(title, subtitle string) {
    // 复杂的 header 渲染逻辑...
}
```

### 新 cmdux 代码

```go
// 使用 cmdux
import "github.com/bagaking/cmdux"
import "github.com/bagaking/cmdux/ui"

app := cmdux.New()
box := ui.NewBox().
    Title("ccmodel").
    Content("AI Model Manager").
    BorderStyle(app.Theme().Primary)
app.Render(box)
```

## 详细迁移步骤

### 1. 替换颜色系统

**原代码:**
```go
Primary.Sprint("重要文本")
Success.Sprint("成功消息")
```

**新代码:**
```go
app.Theme().Primary.Sprint("重要文本")
app.Theme().Success.Sprint("成功消息")
```

### 2. 替换 UI 组件

#### 盒子组件
**原代码:**
```go
func InfoBox(title string, items []string) {
    fmt.Printf("%s %s\n", Accent1.Sprint(Diamond), Bold.Sprint(title))
    for _, item := range items {
        fmt.Printf("%s%s %s\n", Indent, Muted.Sprint(Bullet), item)
    }
}
```

**新代码:**
```go
content := strings.Join(items, "\n")
box := ui.NewBox().
    Title(title).
    Content(content).
    TitleStyle(app.Theme().Accent1).
    ContentStyle(app.Theme().Primary)
app.Render(box)
```

#### 表格组件
**原代码:**
```go
func ModelEntry(index int, name, size, modified string, isActive bool) {
    // 复杂的表格行渲染...
}
```

**新代码:**
```go
table := ui.NewTable().
    Headers("Index", "Name", "Size", "Modified", "Status").
    AddRow("1", "model1", "1.2GB", "2024-01-01", "Active").
    HeaderStyle(app.Theme().Header).
    RowStyle(app.Theme().Primary)
app.Render(table)
```

### 3. 替换动画效果

#### 加载器
**原代码:**
```go
func LoadingAnimation(duration time.Duration, message string) {
    spinner := NewSpinner("dots")
    spinner.Start(message)
    // ...
}
```

**新代码:**
```go
spinner := ux.NewSpinner(ux.SpinnerDots).Color(app.Theme().Primary)
spinner.Start("Loading...")
time.Sleep(2 * time.Second)
spinner.Success("Complete!")
```

#### 进度条
**原代码:**
```go
type ProgressBar struct {
    width, current, total int
    // ...
}
```

**新代码:**
```go
progress := ux.NewProgressBar(30).
    SetTotal(100).
    SetPrefix("Progress").
    Color(app.Theme().Success)
    
for i := 0; i <= 100; i += 10 {
    progress.Update(i)
    time.Sleep(100 * time.Millisecond)
}
progress.Complete("Done!")
```

### 4. 替换交互组件

#### 菜单系统
**原代码:**
```go
func InteractiveMenu(title string, options []string) (int, error) {
    // 复杂的菜单逻辑...
}
```

**新代码:**
```go
menu := ui.NewMenu().
    Title("Choose an option:").
    Options("Option 1", "Option 2", "Option 3").
    TitleStyle(app.Theme().Header).
    SelectedStyle(app.Theme().Selected)
app.Render(menu)
```

#### 用户输入
**新功能 (cmdux 独有):**
```go
// 简单提示
name, _ := input.NewPrompt("Your name?").
    Default("Anonymous").
    Required(true).
    Run()

// 确认提示
confirmed, _ := input.Confirm("Continue?", true)

// 选择菜单
_, choice, _ := input.Select("Choose:", []string{"A", "B", "C"})

// 完整表单
form := input.NewForm("Registration").
    TextField("name", "Name", true).
    TextField("email", "Email", true).
    BooleanField("newsletter", "Subscribe?", true)
results, _ := form.Run()
```

## 迁移清单

- [ ] 更新导入语句
- [ ] 替换颜色变量为主题调用
- [ ] 重构 UI 组件为 cmdux 组件
- [ ] 更新动画和效果
- [ ] 添加交互功能
- [ ] 测试所有迁移的功能
- [ ] 更新文档和示例

## 优势

1. **模块化**: UI/UX 逻辑独立于业务逻辑
2. **重用性**: cmdux 可用于其他项目
3. **可扩展性**: 组件系统易于扩展
4. **一致性**: 统一的 API 设计
5. **主题化**: 内置主题系统
6. **现代化**: 基于流行的设计模式

## 下一步

1. 完成所有组件的迁移
2. 创建自定义主题
3. 添加更多交互组件
4. 优化性能
5. 完善文档