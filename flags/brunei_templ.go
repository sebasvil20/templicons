// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Brunei(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#f4e24e\" d=\"M0 0h30v24H0z\"></path> <path fill=\"#fff\" d=\"M30 15 0 9V4l30 6z\"></path> <path fill=\"#000\" d=\"M30 20 0 14V9l30 6z\"></path> <path fill=\"#bf2a2c\" d=\"M14.19 17.96a4 4 0 0 1-1.8-.67q-.03.01-.03.14.02.29-.3.37c-.2.05-.51-.05-.76-.26a2 2 0 0 0-.82-.45 1.3 1.3 0 0 0-.93.12l-.09.04c-.04 0-.03-.23 0-.3q.1-.16.32-.24l.13-.07-.06-.12q-.14-.27.12-.46c.17-.13.46-.16.71-.06l.14.05q.05.01-.08-.19a.5.5 0 0 1-.12-.42.5.5 0 0 1 .45-.45c.2 0 .31.06.6.33a4.8 4.8 0 0 0 2.85 1.37c.27.03.81.02 1.13-.03a5 5 0 0 0 2.71-1.35c.21-.2.32-.25.49-.25q.21-.01.36.16a.5.5 0 0 1 .12.38l-.08.26q-.16.22.1.1c.25-.1.56-.05.73.1q.21.2.04.47l-.07.14c0 .02.03.03.18.05q.3.07.31.4 0 .13-.02.13l-.16-.06q-.3-.13-.64-.12-.21 0-.42.06-.34.12-.64.38c-.21.19-.35.25-.55.25q-.29 0-.4-.17c-.03-.06-.03-.1-.03-.22v-.16l-.16.1a2 2 0 0 1-.61.33 6 6 0 0 1-2.72.32\"></path> <path fill=\"#bf2a2c\" d=\"M14.54 16.49a4.5 4.5 0 0 1-3.24-1.85 5 5 0 0 1-.82-3.88c.13-.6.4-1.2.85-1.64a4.1 4.1 0 0 0 3.07 5.74c.65.2 1.33 0 1.96-.21a3.9 3.9 0 0 0 2.58-3.28c.11-.8-.02-1.62-.37-2.35.2.1.46.53.61.82.53 1.03.55 2.26.24 3.35a4.6 4.6 0 0 1-4.88 3.3M7.13 9.13c.15-.08.8.1.88.16.08.07.21.38.22.73q.16 0 .35.27c.14.16.34.47.55.53s.37.36.38.74c0 .37.01 2.34.04 2.47s.18.55.22.67c.05.13.07.15-.18.3s-1.39.13-1.6-.14c-.14-.2 0-.28.09-.38s.12-.33.15-.46.1-.27.12-.51-.02-1.32-.06-1.61-.71-.55-.87-.92-.5-1.76-.3-1.85m15.74 0c-.15-.08-.8.1-.88.16-.08.07-.21.38-.22.73q-.16 0-.35.27c-.14.16-.34.47-.55.53s-.37.36-.37.74c-.01.37-.02 2.34-.05 2.47s-.18.55-.22.67c-.05.13-.07.15.18.3s1.39.13 1.6-.14c.14-.2 0-.28-.09-.38s-.12-.33-.15-.46-.1-.27-.12-.51.03-1.32.06-1.61c.04-.3.71-.55.87-.92s.5-1.76.3-1.85\"></path> <path fill=\"#bf2a2c\" d=\"M17.65 7.76c-.5-.24-.94-.53-1.43-.6-.45-.05-.84.28-1.01.47v-.8l.1-.25.29.33.24-.27.4.25s-.32-.46-.3-1.04c0-.14-.35-.25-.78-.26v-.5l.72.33-.41-.55.55-.18-.83-.44q.1-.05.1-.11-.02-.13-.24-.14-.23.01-.25.14.01.08.14.12V5.6c-.42.02-.75.13-.75.26 0 0-.1.8-.33 1.08l.37-.34.25.3.28-.3.13.25v.77c-.18-.2-.56-.5-1-.44-.48.06-.92.35-1.42.59s-1.06.46-1.56.27c.06.65 1.66 1.3 2.16 1.45s1.1-.15 1.33-.23c.22-.07.31.14.25.32-.07.19.01.3.13.4q.03-.2.1-.22v.45a.7.7 0 0 0-.48-.2c.04.12.35.18.27.63-.08.46-.22.98 0 1.39-.27-.12-.37.07-.37.07s.38.13.27.5c-.1.38-.45 1.18.05 1.72-.43-.2-.62-.06-.76-.25 0 .48.7.79 1.03.79h.32c.32 0 1.04-.31 1.05-.79-.14.2-.33.05-.76.25.5-.54.14-1.33.04-1.71s.28-.51.28-.51-.1-.2-.37-.07c.22-.4.07-.93 0-1.39-.09-.45.23-.51.26-.63-.23-.01-.42.14-.5.21v-.48q.08.04.13.24c.11-.1.2-.22.13-.4s.02-.4.24-.32c.23.08.84.38 1.34.23.5-.16 2.1-.8 2.15-1.45-.5.2-1.04-.04-1.55-.27\"></path>")
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
