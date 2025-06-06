// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func PaneuropeanUnion(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#039\" d=\"M0 24h30V0H0z\"></path> <path fill=\"#fc0\" d=\"m14.22 5.08.3-.92-.79-.57h.97l.3-.92.3.92h.97l-.79.57.3.92L15 4.5z\"></path> <path fill=\"#fc0\" d=\"m15 18.67-.3.92h-.97l.79.57-.3.92.78-.57.78.57-.3-.92.79-.57h-.97z\"></path> <path fill=\"#fc0\" d=\"m7 10.67-.3.92h-.97l.79.57-.3.92.78-.57.78.57-.3-.92.79-.57H7.3z\"></path> <path fill=\"#fc0\" d=\"m10.22 6.15.78-.57.78.57-.3-.92.79-.57h-.97l-.3-.92-.3.92h-.97l.79.57z\"></path> <path fill=\"#fc0\" d=\"m8.86 9.08-.3-.92.78-.57h-.97l-.3-.92-.3.92H6.8l.79.57-.3.92.78-.57z\"></path> <path fill=\"#fc0\" d=\"M9.34 15.59h-.97l-.3-.92-.3.92H6.8l.79.57-.3.92.78-.57.79.57-.3-.92z\"></path> <path fill=\"#fc0\" d=\"M12.27 18.52h-.97l-.3-.93-.3.93h-.97l.79.57-.3.92.78-.57.78.57-.3-.92z\"></path> <path fill=\"#fc0\" d=\"m23 10.67.3.92h.97l-.79.57.3.92-.78-.57-.78.57.3-.92-.79-.57h.97z\"></path> <path fill=\"#fc0\" d=\"M19.78 6.15 19 5.58l-.78.57.3-.92-.79-.57h.97l.3-.92.3.92h.97l-.79.57z\"></path> <path fill=\"#fc0\" d=\"m21.14 9.08.3-.92-.78-.57h.97l.3-.92.3.92h.97l-.79.57.3.92-.78-.57z\"></path> <path fill=\"#fc0\" d=\"M20.66 15.59h.97l.3-.92.3.92h.97l-.79.57.3.92-.78-.57-.79.57.3-.92z\"></path> <path fill=\"#fc0\" d=\"M17.73 18.52h.97l.3-.93.3.93h.97l-.79.57.3.92-.78-.57-.78.57.3-.92z\"></path> <path fill=\"#ffcd00\" d=\"M20.89 12.02a5.96 5.96 0 1 1-11.91 0 5.96 5.96 0 0 1 11.91 0\"></path> <path fill=\"#d80000\" d=\"M14 17.9V13H9.06A6 6 0 0 0 14 17.9\"></path> <path fill=\"#d80000\" d=\"M16 17.88A6 6 0 0 0 20.8 13H16z\"></path> <path fill=\"#d80000\" d=\"M20.8 11A6 6 0 0 0 16 6.16V11z\"></path> <path fill=\"#d80000\" d=\"M14 6.14A6 6 0 0 0 9.07 11H14z\"></path>")
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
