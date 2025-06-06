// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func EquatorialGuinea(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#fff\" d=\"M0 8h30v8H0z\"></path> <path fill=\"#579a2c\" d=\"M30 8H0V0h30z\"></path> <path fill=\"#d13626\" d=\"M30 24H0v-8h30z\"></path> <path fill=\"#3272c9\" d=\"M0 24V0l10 12z\"></path> <path fill=\"#000\" d=\"M15.32 9.35h-2.13v4.21c0 .86.49 1 .74 1h.81c.52 0 .58.42.58.42s.06-.42.57-.42h.82c.25 0 .73-.14.73-1v-4.2z\"></path> <path fill=\"#fff\" d=\"M15.32 14.6a.8.8 0 0 0-.58-.22h-.81c-.17 0-.56-.08-.56-.82V9.53h3.9v4.03c0 .74-.4.82-.56.82h-.82q-.4.02-.57.23\"></path> <path fill=\"#000\" d=\"M12.2 12.63s.21.4.5.65c.5.49-.32.88-.32.88s-.26-.17-.34-.36c-.08-.2.21-.78.16-1.17\"></path> <path fill=\"#fff\" d=\"M12.4 13.95a1 1 0 0 1-.2-.22l.07-.37.06-.22q.1.16.24.27c.13.12.12.2.11.22-.02.12-.17.24-.28.32\"></path> <path fill=\"#000\" d=\"M12.44 13.54s0 .7.68.96c.84.34 1.42.2 1.49.22.07.03 0 .71 0 .71s-.93.15-1.75-.23c-.82-.4-.78-.7-.73-.78.04-.08.1-.14.13-.36.04-.2.18-.52.18-.52\"></path> <path fill=\"#fff\" d=\"M14.16 15.28c-.32 0-.79-.04-1.22-.24-.6-.28-.67-.5-.65-.54q.08-.13.14-.37c.12.2.3.41.63.54q.67.27 1.4.23l-.02.37z\"></path> <path fill=\"#000\" d=\"M18.44 12.63s-.22.4-.5.65c-.51.49.31.88.31.88s.26-.17.34-.36c.08-.2-.2-.78-.15-1.17\"></path> <path fill=\"#fff\" d=\"M18.24 13.95a1 1 0 0 0 .19-.22c.01-.05-.04-.25-.07-.37l-.05-.22-.25.27c-.12.12-.11.2-.1.22 0 .12.16.24.28.32\"></path> <path fill=\"#000\" d=\"M18.19 13.54s0 .7-.68.96c-.84.34-1.41.2-1.48.22-.07.03 0 .71 0 .71s.92.15 1.74-.23c.82-.4.78-.7.74-.78s-.1-.14-.14-.36c-.04-.2-.18-.52-.18-.52\"></path> <path fill=\"#fff\" d=\"M16.48 15.28c.31 0 .78-.04 1.21-.24.6-.28.67-.5.66-.54q-.09-.13-.15-.37a1.2 1.2 0 0 1-.62.54q-.69.27-1.4.23l.01.37z\"></path> <path fill=\"#000\" d=\"M16.44 14.98s0 .3-1.12.3-1.13-.3-1.13-.3-.04.48 0 .67.33.3 1.13.3 1.08-.1 1.12-.3 0-.67 0-.67\"></path> <path fill=\"#fff\" d=\"M15.32 15.78c-.9 0-.95-.15-.96-.16l-.01-.29q.28.13.97.13t.96-.13l-.01.29c0 .01-.06.16-.95.16\"></path> <path fill=\"#ffd701\" d=\"m14.8 8 .14.26h.3l-.15.26.15.26h-.3l-.15.26-.15-.26h-.3l.15-.26-.15-.26h.3zm-1.2.14.21.22.3-.07-.1.29.21.22-.3.06-.08.3-.2-.23-.3.07.09-.29-.2-.22.29-.06zm-1.1.38.23.18.28-.12-.04.3.25.17-.28.12-.03.3-.24-.18-.28.13.04-.3-.25-.18.28-.12zM15.9 8l-.14.26h-.3l.14.26-.14.26h.3l.14.26.16-.26h.3l-.16-.26.15-.26h-.3zm1.2.14-.21.22-.3-.07.1.29-.21.22.3.06.08.3.2-.23.3.07-.1-.29.21-.22-.29-.06zm1.1.38-.23.18-.28-.12.03.3-.24.17.28.12.03.3.24-.18.27.13-.03-.3.24-.18-.27-.12z\"></path> <path fill=\"#a36629\" d=\"M15.32 14.34c-.54 0-.62-.2-.44-.61q.28-.64.28-1.66c0-.68-.24-.84-.53-1.1-.28-.26.15-.17.34 0s.27.32.27.32v-.64l.22.07v.48s.32-.27.54-.3c.21-.03.34.14.06.22-.28.09-.6.23-.6.95 0 .71.11 1.46.25 1.73.14.26.14.54-.4.54\"></path> <path fill=\"#009a3b\" d=\"M14.7 11.6c-.28.16-.48-.17-.48-.17s-.26.18-.5.05c-.24-.12-.2-.23 0-.34s.45-.14.45-.14.06-.3.17-.51a.7.7 0 0 1 .3-.3s.22-.48.6-.5c.39-.03.39.37.39.37s.24 0 .29.14c.2-.23.58-.2.67-.1.08.1 0 .3 0 .3s.33-.1.4.02c.07.1 0 .25 0 .25s.26.05.14.34c-.13.3-.52.21-.52.21s-.15.17-.32.15c-.17-.01-.23-.25-.23-.25s-.26.07-.43-.16c-.17.25-.47.1-.51.01-.04.32-.13.49-.42.64\"></path>")
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
