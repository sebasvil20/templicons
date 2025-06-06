// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Olympic(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#fff\" d=\"M0 24h30V0H0z\"></path> <path fill=\"#0081c8\" d=\"M6.59 7.24a2.95 2.95 0 1 0 0 5.9 2.95 2.95 0 0 0 0-5.9M2.4 10.19a4.19 4.19 0 1 1 8.38 0 4.19 4.19 0 0 1-8.38 0\"></path> <path fill=\"#000\" d=\"M15.07 7.24a2.95 2.95 0 1 0 0 5.9 2.95 2.95 0 0 0 0-5.9m-4.2 2.95a4.19 4.19 0 1 1 8.39 0 4.19 4.19 0 0 1-8.38 0\"></path> <path fill=\"#ee334e\" d=\"M23.55 7.24a2.95 2.95 0 1 0 0 5.9 2.95 2.95 0 0 0 0-5.9m-4.2 2.95a4.19 4.19 0 1 1 8.38 0 4.19 4.19 0 0 1-8.37 0\"></path> <path fill=\"#fcb131\" d=\"M10.83 10.86a2.95 2.95 0 1 0 0 5.9 2.95 2.95 0 0 0 0-5.9m-4.2 2.95a4.19 4.19 0 1 1 8.39 0 4.19 4.19 0 0 1-8.38 0\"></path> <path fill=\"#00a651\" d=\"M19.3 10.86a2.95 2.95 0 1 0 0 5.9 2.95 2.95 0 0 0 0-5.9m-4.18 2.95a4.19 4.19 0 1 1 8.37 0 4.19 4.19 0 0 1-8.37 0\"></path> <path fill=\"#0081c8\" d=\"M9.54 10.19c0-.64-.2-1.26-.6-1.77l1-.74a4.2 4.2 0 0 1 0 5.02l-1-.74c.4-.51.6-1.13.6-1.77\"></path> <path fill=\"#000\" d=\"M18.02 10.19a3 3 0 0 0-.6-1.77l1-.74a4.2 4.2 0 0 1 0 5.02l-1-.74c.39-.51.6-1.13.6-1.77m-4.72 2.36c.5.38 1.13.59 1.77.59v1.24c-.9 0-1.8-.3-2.52-.84z\"></path> <path fill=\"#ee334e\" d=\"M21.78 12.55c.5.38 1.13.59 1.76.59v1.24c-.9 0-1.78-.3-2.5-.84z\"></path>")
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
