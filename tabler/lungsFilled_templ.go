// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabler

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func LungsFilled(props ...Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path d=\"M12 3a1 1 0 0 1 1 1v5a2 2 0 0 0 1 1.732v-3.475c0 -1.242 .995 -2.257 2.233 -2.257c.372 0 .738 .094 1.122 .307l.18 .117c1.695 1.23 2.76 3.035 3.773 6.34q .674 2.204 .692 5.06c.016 2.195 -1.657 4.024 -3.843 4.168l-.237 .008c-2.17 0 -3.92 -1.787 -3.92 -3.98v-4.146a4 4 0 0 1 -1.893 -1.112l-.107 -.118l-.107 .118a4 4 0 0 1 -1.892 1.112l-.001 4.146c0 2.193 -1.75 3.98 -3.919 3.98l-.268 -.01c-2.155 -.142 -3.827 -1.971 -3.811 -4.165q .018 -2.858 .692 -5.06c1.011 -3.307 2.076 -5.112 3.822 -6.375l.188 -.117a2.2 2.2 0 0 1 1.064 -.273c1.237 0 2.232 1.015 2.232 2.257l.001 3.475a2 2 0 0 0 .999 -1.732v-5a1 1 0 0 1 1 -1\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = i.Icon("0 0 24 24", props...).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
