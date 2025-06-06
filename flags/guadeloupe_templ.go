// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func Guadeloupe(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#c83029\" d=\"M30 24H0V8h30z\"></path> <path fill=\"#0a2383\" d=\"M30 8H0V0h30z\"></path> <path fill=\"#f6dd4c\" d=\"M14.63 1.9c-.33.53-.37.96-.14 1.7.1.34.22.86.27 1.14l.07.51h.35l.05-.5c.04-.26.14-.79.25-1.15.21-.74.2-1.19-.14-1.7-.12-.18-.29-.34-.36-.34s-.24.14-.35.34m1.6 1.5c-.47.2-.74.75-.9 1.5l-.04.35h.23l.08-.31c.19-.6.74-.83.71-.28q-.01.19-.14.32c.09-.02.27-.1.42-.2.85-.5.52-1.75-.37-1.37m-1.88 2.04h1.31v-.1h-1.31zm.1.44c-.13.14-.34.25-.34.25s.19.02.36 0a1 1 0 0 0 .37-.07c-.03.2-.03.4-.03.4.02.26.05.33.11.46q.1.15.18 0 .09-.18.12-.46a2 2 0 0 0-.05-.4c.12.05.2.06.37.08.18.02.35 0 .35 0s-.25-.11-.36-.27c-.1-.15-.08-.12-.16-.31h-.72c-.09.17-.1.2-.2.32m-.67-2.47c.46.2.73.74.88 1.5.03.13.04.18.05.34h-.23l-.08-.31c-.19-.6-.74-.83-.71-.28q.01.19.14.32a2 2 0 0 1-.43-.2c-.84-.5-.5-1.75.38-1.37M6.75 1.9c-.32.53-.36.96-.14 1.7.1.34.23.86.27 1.14.07.39.06.43.07.51h.36q0-.2.05-.5c.03-.26.13-.79.24-1.15.22-.74.2-1.19-.13-1.7-.13-.18-.29-.34-.36-.34s-.24.14-.36.34m1.6 1.5c-.46.2-.74.75-.89 1.5l-.05.35h.24l.08-.31c.18-.6.73-.83.7-.28q0 .19-.14.32c.1-.02.28-.1.43-.2.85-.5.51-1.75-.37-1.37M6.47 5.45h1.32v-.1H6.47zm.1.44c-.12.14-.33.25-.33.25s.18.02.35 0a1 1 0 0 0 .37-.07l-.03.4c.03.26.06.33.12.46q.1.15.18 0 .1-.18.11-.46a2 2 0 0 0-.05-.4c.12.05.2.06.37.08.19.02.35 0 .35 0s-.24-.11-.35-.27c-.1-.15-.08-.12-.16-.31h-.73c-.08.17-.09.2-.2.32M5.9 3.4c.46.2.74.74.89 1.5.03.13.04.18.05.34H6.6l-.08-.31c-.18-.6-.73-.83-.7-.28q0 .19.14.32a2 2 0 0 1-.43-.2c-.84-.5-.51-1.75.37-1.37m17-1.85c-.08 0-.24.15-.37.34-.33.51-.35.96-.13 1.7.1.36.2.89.24 1.14l.05.51h.36c0-.08 0-.12.07-.51.04-.28.16-.8.27-1.14.22-.74.18-1.17-.14-1.7-.11-.2-.28-.34-.36-.34M21.64 3.4c-.88-.38-1.21.87-.37 1.38.15.08.34.17.43.2a.5.5 0 0 1-.15-.32c-.02-.55.53-.32.71.27.05.15.06.2.08.31h.24l-.05-.35c-.15-.75-.43-1.3-.89-1.5m1.88 1.94H22.2v.1h1.32zm-.1.53c-.11-.11-.12-.15-.2-.32h-.73c-.07.19-.05.16-.16.31s-.35.28-.35.28l.35-.01c.17-.02.25-.03.37-.08a2 2 0 0 0-.05.4q.02.28.11.46.08.15.18 0c.06-.13.1-.2.12-.46 0 0 0-.2-.03-.4a1 1 0 0 0 .37.08h.35s-.2-.12-.34-.26m.67-2.47c.88-.38 1.21.87.37 1.38a2 2 0 0 1-.43.2q.14-.14.15-.32c.02-.55-.53-.32-.71.27l-.08.31h-.24l.05-.35c.15-.75.43-1.3.89-1.5\"></path> <path fill=\"#377e22\" d=\"M7.3 12.05c2.12 1.52-1.1-1.28-1.3-1.44.12.45.68 1.05 1.3 1.44\"></path> <path fill=\"#377e22\" d=\"M11.6 13c-.51-.1-1.25-1.08-1.35-1.7-.01-.23-.22-.2-.13.06.16.68 1.18 1.99.16.98-.24-.22-.58-.76-.77-.86 0 .59 1.26 1.35-.07 1.16-.05-.16-.8-1.18-.7-.9-.05.09 1.01 1.36.34.78-.4-.43-.99-1.24-1.18-1.6-.7-.78.8 1.54.91 1.62-.12.25-2.06-2.82-1.71-1.9.58 1.05.77 1.37-.35-.02-.19-.24-.65-.93-.25-.2.26.5 1.3 1.62 1.88 2.05.22.16.53.26.78.29.01.28 1.04.6.74.74-.92-.65-1.4-.9-2.53-1.17-.33-.23-1.07-.94-1.73-1.34-1.58-.88 1.1.9.3.55-1.48-.55.62.49.8.62-.75-.21-1.95-.46-1.7-.3.18.08 2.73.49 1.64.65-.19.01-.54-.36-.14.04-.3.09-.1.06.06.1 1.17.56-.67.24-.4.85.09-.4 1.32-.54 2 .05-.6-.04-1.94.12-1.69.42.55-.26.96-.52.72.3.12-.52.96-.7 1.69-.42-.21-.11-1.93.68-1.26.55a4 4 0 0 1 1.63-.46c-.37.74-1.61.68-1.2 1.37.1.44.05-.44.23-.49.54-.42.83-.1.3.48.3-.3.4-.45.5.05-.19-.65.33-.73.98-.55-.7.3-.72.74-.68 1.33.02-.71.62-1.45 1.4-1.15-.4-.3.23-.23.24-.35-.04-.11-.82.13-1.05-.04-1.15-.02-1.06-.04-.05-.24.71-.2.07-.28-.76.01-.02-.67 1.4-.15 1.58-.4-.13-.1-.74-.34-.63-.47.46.24.95.45 1.05.4.27-.17-1.58-.44-.59-.57 1 .25-.24-.28-.43-.48.31-.12.54.4 1.12.37.5.02-.29-.28-.5-.37.15-.02 1.5.54.8.16m-1.42.79c-.1-.05-.43-.04-.22-.15q.08.07.22.14m-2.8-1.24c.1-.04.16 0 .37.07-.13-.01-.27.03-.37-.07m.03.23c.6.4-.67.06-.15.02zm.36.04c-.22-.33.47 0 .63.27-.31-.1-.4-.1-.63-.27m1.05.54.4.23c-.27.04-1.01-.4-.4-.22m.93.22c-.14.13-.26-.04-.4-.1zm-.78-.4c-.2 0-.38-.23-.54-.3.37.14.6.26.54.3m-2.12-.51c.14 0 .3-.18.4-.05 0 .1-.3.2-.4.05m1.26.69q-.06.01-.4-.16c.1 0 .43.05.4.16m2.03 0a2 2 0 0 1-.71-.47c.1.15 1.18.27.7.46m-.26-.55c.25.21-.1.04-.3 0 .07 0 .2-.02.3 0\"></path> <path fill=\"#377e22\" d=\"M9.87 12.63c.17-.04-.98-1.33-.97-1.32.05.26.89 1.52.97 1.32m-3.32-.4c-.82-.42-1.4.74-1.04.55.2-.61 1.22-.4 1.04-.54m3.87 3c-.28-.12-.79 1.1-.37.54.04-.2.72-.57.37-.53m13.65 3.02h-.01c-.17.4.33-.72 0 0\"></path> <path fill=\"#377e22\" d=\"M20.93 17.23c-5.1-1.4 1.99.82 3.13 1.01.17-.5-2.45-.67-3.13-1.01m.18-.28c-5.14-1.2 2.01.76 3.17.91.15-.52-2.48-.59-3.17-.9\"></path> <path fill=\"#377e22\" d=\"M24.28 17.86c-.16.42.3-.72 0 0m-.51.8c-.2.4.36-.7 0 0\"></path> <path fill=\"#377e22\" d=\"M20.7 17.48c-5.03-1.66 1.93.93 3.06 1.18.2-.5-2.4-.8-3.07-1.18m2.8 1.6a90 90 0 0 1-4.96-1.85c-.4.25 4.23 1.94 4.73 2.24.22-.25-.02-.24-.3-.37.19.06.55.26.52.02.09-.14.3-.5.02-.04\"></path> <path fill=\"#377e22\" d=\"M23.5 18.85c.06-.28-4.09-1.4-4.8-1.83-.93-.12 5.49 2.57 4.8 1.83m-.22.62c-.23.38.42-.66 0 0m-3.18-1.15c-4.74-2.34 1.79 1.18 2.88 1.59.27-.47-2.28-1.13-2.88-1.6m2.88 1.6c-.24.37.46-.64 0 0\"></path> <path fill=\"#f6dd4c\" d=\"m15.94 19.36.26.19q.58.41.56 1l-.01.3.12-.3c.14-.32.12-.8-.04-1.16l-.1-.2.3.13q.65.27.76.85l.05.3.06-.31c.06-.34-.06-.8-.29-1.12l-.13-.19.32.06c.45.1.78.34.91.69l.12.29-.01-.33a1.6 1.6 0 0 0-.51-1.04l-.17-.15h.32c.45-.01.84.16 1.04.47l.17.26-.08-.32a1.6 1.6 0 0 0-.71-.9l-.2-.12.31-.07c.45-.1.86-.02 1.12.25l.22.22-.14-.3a1.6 1.6 0 0 0-.9-.74l-.21-.07.3-.13c.4-.2.82-.2 1.14.01l.26.17-.2-.26a1.6 1.6 0 0 0-1.03-.54l-.23-.02.27-.2c.36-.27.77-.35 1.12-.22l.29.1-.25-.2a1.6 1.6 0 0 0-1.12-.32l-.23.03.22-.24c.3-.35.68-.52 1.05-.46l.3.05-.28-.15a1.6 1.6 0 0 0-1.16-.08l-.22.07.16-.28q.35-.62.94-.66l.3-.02-.31-.09a1.6 1.6 0 0 0-1.15.16l-.2.12.1-.31q.21-.68.78-.84l.3-.08-.33-.02a1.6 1.6 0 0 0-1.1.4l-.16.15.03-.32c.04-.46.25-.82.58-.98l.28-.15-.33.05c-.34.04-.75.3-.98.61l-.13.19-.04-.32c-.05-.45.08-.85.37-1.08l.24-.2-.31.11c-.33.11-.67.45-.83.8l-.1.21-.1-.3c-.15-.43-.1-.85.14-1.13l.19-.25-.28.18c-.3.17-.57.58-.65.96l-.04.22-.17-.28c-.23-.39-.28-.8-.1-1.13l.14-.28-.24.23c-.25.23-.43.68-.43 1.07v.23l-.22-.24c-.31-.33-.44-.73-.34-1.1l.07-.29-.18.27c-.2.29-.28.76-.2 1.14l.05.22-.26-.18c-.38-.26-.58-.63-.56-1l.01-.3-.12.3c-.14.31-.12.8.04 1.15l.1.21-.3-.13c-.42-.17-.7-.49-.76-.86l-.05-.3-.06.32c-.06.34.06.8.29 1.12l.13.19-.31-.07c-.46-.08-.8-.33-.93-.68l-.1-.29v.33c.01.34.22.78.51 1.04l.17.15h-.32c-.45.01-.84-.16-1.04-.48l-.17-.25.08.31c.08.34.38.72.71.91l.2.12-.31.07c-.45.1-.86.01-1.12-.25l-.22-.22.14.3c.15.3.52.62.9.74l.21.07-.3.13c-.4.2-.82.19-1.14-.02l-.26-.16.2.26c.21.27.64.5 1.03.54l.23.02-.27.2c-.36.27-.77.35-1.12.22l-.29-.11.25.21c.26.23.73.36 1.12.32l.23-.03-.22.24c-.3.35-.68.51-1.05.46l-.3-.05.28.15c.3.17.79.2 1.16.08l.22-.07-.16.28c-.22.4-.56.64-.94.66l-.3.02.31.09c.33.1.81.03 1.15-.16l.2-.12-.1.3q-.21.69-.78.85l-.3.08.33.02c.35.03.8-.13 1.1-.4l.16-.15-.03.32c-.04.46-.25.82-.59.99l-.27.14.32-.05c.35-.04.76-.3.99-.61l.13-.19.04.32c.06.45-.08.85-.37 1.09l-.24.2.31-.12c.33-.11.67-.45.83-.8l.1-.21.1.3c.15.43.1.85-.14 1.14l-.19.24.28-.18c.3-.17.57-.57.65-.96l.04-.22.17.28c.23.39.28.8.1 1.14l-.13.27.23-.23c.26-.23.44-.68.44-1.07v-.23l.21.24c.31.33.44.73.34 1.1l-.07.3.18-.28c.2-.28.28-.76.2-1.14z\"></path>")
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
