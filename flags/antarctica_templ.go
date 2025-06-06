// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Antarctica(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#265fb5\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fff\" d=\"M23.3 12.79c-.05-.44.67-1.55.1-1.47-.71.1-1.27-.39-1.5-.38.76-1.08.2-1.36.02-2.06-.3-1.23.66-1.84-1.06-2.1-1.11-.17-1.89-.06-2.22-1.33-.5.34-1.08.43-1.64.19-.35-.16-.46-.45-.8-.6-.25-.1-.78-.1-1.06-.15-.79-.12-1.19.22-1.89.1.07.11.08.24.13.35-.44.01-.7.3-.75.7-.63-.06-.79.77-.72 1.29l-.06.02.02.02c-.13 1.38-.69 2.06-2.1 2.28-.93.14-1.75-.47-2.55-.97a4 4 0 0 1-.65-.53c-.24-.26-.4-.63-.63-.82-.11-.08-.34-.16-.3.02-.15.91.52 2.05 1.28 2.49-1.05.48.27.77.46.94l.46.65c.07.44-.24.4-.27.75-.03.37-.03.66.1 1.03l-.03.03c0 .3-.22.19-.51.12.04.9.9.98 1.03 1.15.46.57.3 1.4.84 1.92 1.33 1.3 2.84.53 4.34.75.84.12 1.56.83 2.32.03.9.98-.7 2.15-.3 2.48q.16.12.34.14l-.07.18c1.08.1 2.21.36 3.22-.17 0 .16.1.35.11.44.25-.3.44-.28.76-.44.4-.2 1.08-.43 1.35-.82.3-.47-.18-1.12.8-.92.07-.34.1-.3.29-.53-.27-.67.53-1.13.77-1.61.05-.1.05-.56.12-.73.12-.32.46-.5.56-.91.18-.75-.25-.88-.31-1.53\"></path>")
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
