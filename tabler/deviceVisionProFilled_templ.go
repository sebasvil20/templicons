// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabler

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func DeviceVisionProFilled(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path d=\"M12 6q 1.74 0 3.342 .106q 1.619 .107 2.973 .448q 1.388 .345 2.436 1.05a4.9 4.9 0 0 1 1.665 1.916c.397 .801 .584 1.769 .584 2.91c0 1.156 -.222 2.208 -.673 3.14c-.45 .934 -1.073 1.685 -1.868 2.236a4.7 4.7 0 0 1 -2.73 .839q -.932 .001 -1.703 -.263a7 7 0 0 1 -1.374 -.644a20 20 0 0 1 -1.107 -.736a8 8 0 0 0 -.901 -.567a1.4 1.4 0 0 0 -.643 -.174c-.209 0 -.426 .057 -.658 .18q -.42 .226 -.893 .564a20 20 0 0 1 -1.105 .733a6.8 6.8 0 0 1 -1.366 .642a5.2 5.2 0 0 1 -1.688 .264a4.7 4.7 0 0 1 -2.75 -.838c-.794 -.55 -1.418 -1.302 -1.868 -2.234q -.675 -1.407 -.673 -3.14c-.005 -1.135 .182 -2.105 .577 -2.9a4.9 4.9 0 0 1 1.673 -1.926c.699 -.47 1.511 -.816 2.442 -1.049a17 17 0 0 1 2.968 -.447q 1.599 -.11 3.34 -.11\"></path>")
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
