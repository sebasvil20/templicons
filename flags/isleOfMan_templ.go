// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func IsleOfMan(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#bf2c30\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fff\" d=\"M22.79 7.02c-.07-.2-.23-.15-.36-.22-.18-.08-.43-.2-1.41.06l-.82-.08c-.35-.33-.82 0-.73.33l.07.37-.13.56a6 6 0 0 0-1.1 2.82c0 .08 0 .1.06.4.01.07-.08.04-.15.03-.79-.46-1.36-.68-2.13-1.06q-.04-.05-.05-.2c.03-.37-.62-1.93-.7-2.14a3 3 0 0 1-.48-1c-.07-.89-1.28-1.17-1.81-.65-.14.2-.2.17-.34.27-.79.1-1.54.44-2.15.71-.68.34-1.1.71-1.7.2L7.69 6.3c-.24-.19-.37-.09-.42-.03-.13.15 0 .26 0 .41.01.2.04.46.76 1.17l.33.74c-.11.46.42.7.66.45l.28-.24.57-.17c1.04.07 2.36-.1 3.03-.47.07-.04.08-.05.32-.25.06-.04.08.05.1.12 0 .89.09 1.47.13 2.3.04.03-.15.2-.25.24-.38.3-1.3 1.45-1.43 1.6-.14.32-.39.69-.63.91-.75.5-.4 1.68.33 1.87.24.03.25.08.4.15.49.62 1.17 1.1 1.71 1.48.64.4 1.17.57 1.03 1.34l-.4 1.58c-.04.32.1.36.21.37.18 0 .2-.14.33-.2.16-.11.38-.27.65-1.24l.48-.65c.47-.13.41-.7.07-.78l-.36-.13-.43-.4a6 6 0 0 0-1.94-2.34l-.37-.15c-.07-.03 0-.1.05-.14.79-.44 1.26-.8 1.97-1.27.05-.05.22.05.32.1.44.18 1.93.39 2.14.41q.57-.05 1.12.1c.75.39 1.7-.49 1.47-1.22-.1-.22-.05-.25-.07-.42.46-1.03.4-2.32.53-3.27.23-.54 1.7-.73 2.18-.9.28-.12.26-.28.24-.34\"></path>")
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
