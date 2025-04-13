// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabler

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func BlenderFilled(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path d=\"M15 14c.899 0 1.728 .296 2.396 .797l-.12 -.086q .101 .069 .196 .144l-.076 -.058q .07 .053 .14 .109l-.064 -.05q .072 .055 .14 .115l-.077 -.065a4 4 0 0 1 .565 .566l-.065 -.078q .061 .072 .12 .147l-.055 -.07q .057 .071 .111 .144l-.056 -.074q .066 .084 .127 .172l-.07 -.098q .058 .078 .112 .16l-.042 -.062q .054 .078 .104 .157l-.062 -.096q .06 .09 .117 .185l-.055 -.089l.09 .15l-.035 -.061a4 4 0 0 1 .338 .727l-.055 -.15q .039 .1 .073 .201l-.018 -.051q .033 .093 .06 .188l-.042 -.137q .033 .1 .061 .204l-.019 -.067q .027 .091 .05 .184l-.03 -.117q .03 .114 .055 .23l-.025 -.113q .026 .107 .045 .215l-.02 -.101q .021 .1 .037 .201l-.017 -.1q .018 .095 .03 .19l-.013 -.09a4 4 0 0 1 .026 .196l.017 .205l.006 .226v2a2 2 0 0 1 -2 2h-8a2 2 0 0 1 -2 -2v-2q 0 -.273 .036 -.537l-.013 .105a4 4 0 0 1 .026 -.194l-.013 .09a4 4 0 0 1 .03 -.191l-.017 .1q .015 -.1 .037 -.2l-.02 .1q .02 -.108 .046 -.215l-.026 .114q .024 -.117 .056 -.23l-.03 .116a4 4 0 0 1 .049 -.184l-.02 .067q .03 -.104 .062 -.204l-.042 .137q .027 -.096 .06 -.188l-.018 .051a4 4 0 0 1 .32 -.717l-.064 .117q .048 -.09 .1 -.178l-.035 .06l.09 -.15l-.055 .09q .055 -.094 .117 -.185l-.062 .096q .05 -.08 .104 -.157l-.042 .061q .054 -.08 .113 -.16l-.071 .1q .06 -.09 .127 -.173l-.056 .074q .054 -.073 .111 -.143l-.055 .069a4 4 0 0 1 .542 -.57l-.072 .064q .073 -.066 .15 -.129l-.078 .065a4 4 0 0 1 .14 -.116l-.062 .051q .068 -.056 .139 -.11l-.076 .06q .096 -.076 .195 -.145l-.12 .086q .07 -.053 .142 -.101l-.022 .015q .072 -.05 .145 -.096a4 4 0 0 1 2.132 -.615zm-2 3a1 1 0 0 0 -.993 .883l-.007 .127a1 1 0 0 0 1.993 .117l.007 -.127a1 1 0 0 0 -1 -1m1 -15a1 1 0 0 1 1 .999l1.802 .001a2 2 0 0 1 1.968 2.358l-1.49 8.191a5 5 0 0 0 -2.28 -.549h-4c-.822 0 -1.597 .198 -2.28 .55l-.464 -2.551l-2.256 .001a2 2 0 0 1 -2 -2v-4a2 2 0 0 1 2 -2l5 -.001l.007 -.116a1 1 0 0 1 .993 -.883zm-6.835 2.999l-1.165 .001v4l1.893 -.001z\"></path>")
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
