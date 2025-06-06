// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Kosovo(props ...i.Props) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#2b48a1\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#c8a75e\" d=\"M15.16 8.2c-.06 1.36 1.36.39 1.83 1.49.49.44.64 1.52 1.73 1.46-.16.78-.33.97.56 1.15.45-.26.47.23.77.39.24-.07.73-.3.86.19-.37.91-.73 1.95-1.3 2.44-1.15-.03 1.32 1.15-.43.92-.24-.59-.68.28-1.04.36-.1 1.02-.7 1.25-1.15.17-3.9.52-1.32 2.44-3.4 2.62-.62-5.26-1.7-2.35-2.5-4.61-.12-.52-.58-.74-.93-1.1.25-.68.08-.87-.58-1.14-.07-.62.7-.87 1.2-.48.44-.13.6-.62 1.23-.65 1.03.04.74-.26.88-.98.52-.41.87-.23 1-.93.49-.17-.87-.71-.12-1.02.46-.08 1.06-.8 1.39-.28\"></path> <path fill=\"#fff\" d=\"m16.71 5.47-.28-.85-.27.85h-.9l.73.53-.28.86.72-.53.73.53-.28-.86.73-.53zm2.83.5-.28-.85-.27.85h-.9l.73.53-.28.85.72-.53.73.53-.28-.85.73-.53zm2.7.98-.27-.85-.28.85h-.9l.73.53-.28.86.73-.53.72.53-.28-.86.73-.53zm-8.4-1.48-.28-.85-.27.85h-.9l.72.53-.27.86.72-.53.73.53L14 6l.73-.53zm-2.83.5-.27-.85-.28.85h-.9l.73.53-.28.85.73-.53.72.53-.28-.85.73-.53zm-2.7.98-.28-.85-.27.85h-.9l.73.53-.28.86.72-.53.73.53-.28-.86.73-.53z\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = i.FlagIcon("0 0 30 24", props...).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
