// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Australia(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#061b65\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fff\" d=\"m5.84 14.74.53 1.67 1.64-.62-.97 1.45 1.5.9-1.74.15.24 1.73-1.2-1.27-1.2 1.27.23-1.73-1.74-.15 1.5-.9-.96-1.45 1.63.62zm16.27 3.01.3.91.88-.34-.52.8.82.48-.95.08.13.95-.66-.7-.65.7.13-.95-.95-.08.82-.48-.53-.8.9.34zm-2.94-8.62.29.9.89-.33-.53.79.82.48-.95.09.13.94-.65-.69-.66.69.13-.94-.95-.09.82-.48-.53-.8.9.34zm3.26-6 .3.9.88-.33-.53.79.82.48-.95.09.14.94-.66-.69-.66.69.14-.94-.95-.08.82-.5-.53-.78.89.33zm4.7 3.69.3.9.88-.33-.52.79.82.48-.95.09.13.94-.66-.69-.65.69.13-.94-.95-.09.82-.48-.53-.8.9.34zM24.74 12l.23.56.61.05-.47.4.15.59-.52-.33-.52.33.15-.6-.47-.39.61-.04zM5.5 9.77V12h4V9.77L12.54 12H15v-1.24L11.24 8H15V4h-2.74L15 2V0h-1.44L9.5 2.98V0h-4v2.8L1.75 0H0v1.9L2.99 4H0v4h3.76L0 10.76V12h2.46z\"></path> <path fill=\"#d22d32\" d=\"m0 1 4.09 3h1.36L0 0zm0 11 5.45-4v1l-4.09 3z\"></path> <path fill=\"#d22d32\" d=\"M8.5 12h-2V7H0V5h6.5V0h2v5H15v2H8.5z\"></path> <path fill=\"#d22d32\" d=\"M15 11.67 10 8v1l4.1 3h.9zM15 0h-.25L9.3 4h1.37L15 .82z\"></path>")
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
