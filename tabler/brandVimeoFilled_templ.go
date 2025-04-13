// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabler

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func BrandVimeoFilled(props ...Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path d=\"M18.5 3c2.38 0 3.985 2.174 3.486 5.164c-.535 3.21 -2.25 6.074 -4.808 8.675c-1.277 1.298 -2.211 2.061 -4.112 3.485c-2.323 1.597 -4.408 .365 -5.47 -1.897c-.292 -.618 -.586 -1.724 -1.248 -4.477l-.03 -.126c-.483 -2.01 -.819 -3.319 -.982 -3.878l-.016 -.052l-.031 .013l-.13 .06l-.137 .07a4 4 0 0 0 -.43 .269a1 1 0 0 1 -1.3 -.099l-1 -1a1 1 0 0 1 -.124 -1.262a20 20 0 0 1 1.918 -2.382c.98 -1.037 1.955 -1.816 2.928 -2.233c.5 -.214 .996 -.33 1.486 -.33c2.237 0 3.02 1.588 3.567 4.963c.03 .183 .057 .359 .112 .709c.123 .784 .197 1.198 .292 1.588c.292 1.185 .528 1.984 .735 2.483l-.016 -.039l.096 -.107c.354 -.411 .757 -1.014 1.172 -1.771l.157 -.291c.391 -.745 .505 -1.528 .363 -1.9c-.028 -.073 .007 -.065 -.456 .218a1 1 0 0 1 -1.51 -1.013c.496 -3.053 2.745 -4.84 5.488 -4.84\"></path>")
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
