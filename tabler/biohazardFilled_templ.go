// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabler

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func BiohazardFilled(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path d=\"M16.41 2.072a6.25 6.25 0 0 1 1.514 6.387l-.051 .137l.223 .044q .325 .072 .645 .18l.318 .117l.25 .105c2.155 .97 3.572 3.067 3.681 5.483v.217a1.5 1.5 0 1 1 -3 -.003l.002 -.145a3.25 3.25 0 0 0 -4.412 -2.886l-.091 .037l.004 .038l.007 .217a3.5 3.5 0 0 1 -1.817 3.07l-.16 .082l.014 .11c.082 .511 .285 .997 .595 1.416l.14 .175a3.25 3.25 0 0 0 2.27 1.136l.203 .006a1.5 1.5 0 0 1 0 3a6.25 6.25 0 0 1 -4.575 -1.991l-.177 -.199l-.078 .092a6.3 6.3 0 0 1 -3.921 2.054l-.273 .028l-.259 .016h-.217a1.5 1.5 0 1 1 .003 -3l.145 .002a3.25 3.25 0 0 0 3.074 -2.82l.003 -.03l-.161 -.083a3.5 3.5 0 0 1 -1.804 -2.883l-.005 -.195l.006 -.191l.003 -.043l-.075 -.032a3.25 3.25 0 0 0 -2.398 .008l-.191 .084a3.25 3.25 0 0 0 -1.85 2.933a1.5 1.5 0 0 1 -3 0a6.25 6.25 0 0 1 5.036 -6.13l.077 -.014l-.05 -.143l-.08 -.26l-.066 -.25a6.27 6.27 0 0 1 1.47 -5.678l.163 -.172a1.5 1.5 0 1 1 2.171 2.07l-.137 .143a3.25 3.25 0 0 0 .386 4.723l.084 .062l.05 -.034a3.5 3.5 0 0 1 1.673 -.555l.228 -.007c.683 0 1.336 .197 1.894 .556l.048 .033l.067 -.048a3.25 3.25 0 0 0 1.111 -1.669l.05 -.2a3.25 3.25 0 0 0 -.74 -2.828l-.141 -.15a1.5 1.5 0 1 1 2.12 -2.122\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = i.FilledIcon("0 0 24 24", props...).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
