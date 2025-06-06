// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Lebanon(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#fff\" d=\"M0 7h30v10H0z\"></path> <path fill=\"#da352e\" d=\"M30 7H0V0h30zm0 17H0v-7h30z\"></path> <path fill=\"#4aa25a\" d=\"M14.76 7.5c-.13.51-.94.76-1 1.18.68-.01-.47.47-.61.65.82.46-.14.25-.29.73-.12.3-1.54.66-.85.9.3-.13.8-.19.65.31-.14.54-1.98.94-1.34 1.45.23-.02.4-.21.6-.3.44-.1.38.53 0 .57-1.83 1.1-1.68 1.67.31.59 1.32.48-1.68.73-1.24 1.2.91.35 1.97-1.1 2.69-.35.68.93-2.18 2.12-1.4 1.99.31-.35.9-.11 1.22-.46s1.06-.46 1.37-.16c.1.17.52.6.65.45.09-.84 1.08.38 1.34.25.06-.32-.46-.4-.61-.65-.48-.3-.8-1.95.04-1.79.57.41 1.27.12 1.9.29.29.07 1.39 1.04 1.13.26-.29-.53-1.76-.68-1.5-1.3.24-.29 1.54.63 1-.1-.5-.29-.8-.96-1.45-.87-.12-.42.68-.17.93-.27.5.06.14-.35-.16-.32-1.04-.18-1.01-.6.04-.32.33-.36-2.65-1.51-1.68-1.73 1.01.38.91.14.11-.35-.14-.13.13-.18.22-.13.24-.74-2.03-.63-2.07-1.72\"></path>")
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
