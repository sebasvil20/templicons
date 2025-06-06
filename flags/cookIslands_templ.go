// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func CookIslands(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#071b65\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fff\" d=\"M5.5 9.77V12h4V9.77L12.54 12H15v-1.24L11.24 8H15V4h-2.74L15 2V0h-1.44L9.5 2.98V0h-4v2.8L1.75 0H0v1.9L2.99 4H0v4h3.76L0 10.76V12h2.46z\"></path> <path fill=\"#b92932\" d=\"m0 1 4.09 3h1.36L0 0zm0 11 5.45-4v1l-4.09 3z\"></path> <path fill=\"#b92932\" d=\"M8.5 12h-2V7H0V5h6.5V0h2v5H15v2H8.5z\"></path> <path fill=\"#b92932\" d=\"M15 11.67 10 8v1l4.1 3h.9zM15 0h-.25L9.3 4h1.37L15 .82z\"></path> <path fill=\"#fff\" d=\"m22.91 7.27.67-.49h-.83L22.5 6l-.25.78h-.83l.67.49-.26.78.67-.48.67.48zm2.1.6.8-.17-.74-.33.08-.82-.55.61-.75-.33.41.71-.55.61.8-.17.42.71zm1.68 1.45.8.17-.54-.61.4-.72-.74.34-.56-.61.1.82-.76.34.8.16.1.82zm.94 1.97.67.48-.26-.78.66-.5-.82.01-.26-.78-.25.79h-.82l.67.48-.25.78zm.07 2.2.41.72.08-.82.81-.18-.75-.33.08-.82-.55.61-.75-.33.41.71-.55.62zm-.85 2 .08.83.41-.71.81.17-.55-.62.41-.7-.75.33-.55-.62.08.82-.75.34zm-1.6 1.55-.24.78.66-.48.67.48-.26-.78.66-.49h-.82l-.26-.78-.25.79h-.82zm-2.06.72-.55.62.8-.18.42.71.08-.82.8-.17-.75-.33.09-.82-.55.61-.75-.33zM20 7.87l-.8-.17.74-.33-.08-.82.55.61.75-.33-.41.71.55.61-.8-.17-.42.71zm-1.7 1.45-.8.17.54-.61-.4-.72.74.34.56-.61-.1.82.76.34-.8.16-.1.82zm-.94 1.97-.67.48.26-.78-.67-.5.83.01.26-.78.25.79h.82l-.67.48.25.78zm-.07 2.2-.41.72-.09-.82-.8-.18.75-.33-.08-.82.55.61.75-.33-.41.71.55.62zm.85 2-.08.83-.41-.71-.81.17.55-.62-.41-.7.75.33.55-.62-.08.82.75.34zm1.6 1.55.24.78-.66-.48-.67.48.26-.78-.66-.49h.82l.26-.78.25.79h.82zm2.06.72.55.62-.8-.18-.42.71-.08-.82-.8-.17.74-.33-.08-.82.55.61.75-.33z\"></path>")
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
