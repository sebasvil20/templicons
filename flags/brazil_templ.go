// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Brazil(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#459a45\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fedf00\" d=\"M2.47 12 15 20l12.53-8L15 4z\"></path> <path fill=\"#002776\" d=\"M14.9 16.85a5 5 0 1 0 0-10 5 5 0 0 0 0 10\"></path> <path fill=\"#cbe9d4\" d=\"M10.49 9.5a5 5 0 0 0-.52 1.56c3.09-.22 7.29 1.47 9.09 3.56a5 5 0 0 0 .68-1.55c-2.22-2.18-6.13-3.59-9.25-3.57\"></path> <path fill=\"#88c9f9\" d=\"M10.27 12.09h.78v.77h-.78zm.78 1.55h.77v.77h-.77z\"></path> <path fill=\"#55acee\" d=\"M12.6 12.09h.77v.77h-.77zm1.54.77h.78v.78h-.78zm3.1 1.55h.78v.77h-.78zm-2.32.77h.77v.78h-.77zm2.32-4.64h.78v.77h-.78z\"></path> <path fill=\"#3b88c3\" d=\"M15.7 13.64h.77v.77h-.78z\"></path>")
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
