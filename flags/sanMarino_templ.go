// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func SanMarino(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#fff\" d=\"M30 24H0V0h30z\"></path> <path fill=\"#75b6e1\" d=\"M30 24H0V12h30z\"></path> <path fill=\"#f4900c\" d=\"M15.05 10.04c-2.15 0-2.35-.85 0-.85s2.16.85 0 .85\"></path> <path fill=\"#ffac33\" d=\"M16.71 9.8c-.54 0-.97.32-1.35.52-.11-.13-.3-.12-.3-.12s-.2-.01-.31.12c-.39-.2-.82-.52-1.36-.52s-1.81.5-1.81 3.1c0 3.4 3.47 4.48 3.47 4.48s3.48-1.08 3.48-4.49c0-2.58-1.28-3.1-1.82-3.1m-4.15-1.14c.35.06.8.68.8.96.32-.35 1.7-.37 1.7-.37s1.36.02 1.68.37c0-.28.46-.9.8-.96-.06-.14-.26-.1-.26-.1s.43-.7.26-1.04-.81-.32-1.34-.38-.62-.28-.62-.28l.14-.16s-.27-.05-.21-.23c.05-.17-.17-.16-.17-.16s.17-.16.1-.3c-.07-.15-.22-.2-.22-.2s-.1-.2 0-.28c.1-.06.32.2.32-.06s-.25-.04-.32-.14c-.07-.09.12-.33-.17-.33s-.1.24-.16.34c-.07.09-.32-.13-.32.13 0 .25.22 0 .32.07.1.06 0 .28 0 .28s-.16.04-.23.19c-.06.14.1.3.1.3s-.22-.01-.17.16c.06.18-.2.23-.2.23l.14.16s-.1.23-.63.29c-.52.05-1.17.04-1.34.37-.17.34.26 1.05.26 1.05s-.2-.05-.26.09m3.66-1.22c.47.08 1.23.24.33.82.01-.2-.1-.37-.1-.37s-.22.03-.4.2c-.08-.27-.2-.36-.2-.36s-.27.08-.35.28l-.13-.22a1 1 0 0 0 .13-.5l-.02-.12q.36.2.74.27m-2.33 0q.37-.06.73-.27l-.01.12q0 .31.12.5l-.12.23c-.08-.2-.35-.29-.35-.29s-.13.09-.2.35c-.18-.16-.4-.19-.4-.19s-.12.17-.1.37c-.9-.58-.14-.74.33-.82\"></path> <path fill=\"#a0cfeb\" d=\"M17.61 13.23c0-2-.6-2.23-.6-2.23s-.28.34-.7.2a5 5 0 0 1-.98-.57s-.1.12-.3.13a.4.4 0 0 1-.28-.13s-.58.44-1 .57c-.4.14-.7-.2-.7-.2s-.59.23-.59 2.23c0 1.89 2.31 3.67 2.57 3.85v.02h.02v-.02c.25-.18 2.56-1.96 2.56-3.85\"></path> <path fill=\"#ccd6dd\" d=\"M14.44 15.13c.14-.39.22-.64.22-.87v-.98s-.22-.1-.22-.23v-.36l.22-.08v.24l.23-.08v-.21h.16v.22l.2.04v-.2l.25.07v.34s-.2.2-.2.25v.98c0 .1.3.87.3.87zm.58-2.76s-.06-.25-.22-.47c-.15-.21-.1-.61.24-.62.62 0 .2.55.2.55s-.13-.28-.19-.2.15.4.06.57zm-1.4.24s-.06-.25-.21-.47c-.16-.21-.1-.61.23-.61.62 0 .2.54.2.54s-.13-.28-.19-.2.15.4.06.57zm2.77 0s-.06-.25-.21-.47-.1-.61.23-.61c.63 0 .21.54.21.54s-.14-.28-.2-.2.15.4.06.57zm-3.24 2.46q.11-.33.12-.57v-.98s-.23-.1-.23-.23v-.36l.23-.08v.24l.22-.08v-.21h.16v.23l.2.04v-.2l.25.06v.34s-.2.2-.2.25v.98c0 .1.3.87.3.87s-1.08-.2-1.05-.3m2.67.3c.14-.39.22-.64.22-.87v-.98s-.22-.1-.22-.23v-.36l.22-.08v.24l.22-.08v-.21h.17v.23l.2.04v-.2l.24.06v.34s-.2.2-.2.25v.98c0 .06.12.4.21.63.05.14-1.06.24-1.06.24\"></path> <path fill=\"#99aab5\" d=\"M13.62 12.8h-.13v.21l-.22.08v-.24l-.22.08v.37c0 .12.22.22.22.22v.98q0 .24-.12.57c-.01.04.22.11.47.18zm1.39-.24h-.13v.21l-.22.08v-.24l-.23.08v.36c0 .13.23.23.23.23v.98q0 .23-.12.57c-.02.04.22.11.47.17zm1.42.28-.18-.04v.2l-.22.08v-.24l-.22.08v.37c0 .12.22.22.22.22v.98q0 .23-.12.57c0 .04.22.11.52.22z\"></path> <path fill=\"#8bb273\" d=\"M15.92 14.99s-.29-.36-.87-.36a1 1 0 0 0-.83.37s-.51-.2-1.1.08a9 9 0 0 0 1.9 2v.02h.02v-.02a9 9 0 0 0 1.93-2c-.64-.27-1.05-.09-1.05-.09\"></path> <path fill=\"#658d5c\" d=\"M15.81 18.86c-.4-.62-.74-1.33-1.7-1.76-.95-.43-1.1-.55-1.38-1a3 3 0 0 0-.82-.73s.01-.41-.17-.68-.16-.6-.16-.6-.22.11-.17-.3c.05-.43-.22-.7-.22-.7s.35-.45.05-1.02c.34-.56.18-1.19.18-1.19s.3-.52 0-1.08c-.92.37-.98 1.57-.98 1.57l-.53-.74s-.44.28-.33.8c.11.53.09.64.09.64l-.5-.4s-.36.57.14 1.02.58.6.58.6-.25-.14-.25.07.05.31.05.31l-.65-.52s-.2.71.26 1.11.74.55.74.55-.87-.24-.86-.18.18.5.62.74c.45.25.78.34.78.34l-.9.16s.22.62.75.62c.52 0 1.02-.16 1.02-.16s-.37.3-.48.52c-.12.23-.43.25-.43.25s.29.3.83.2.6-.8 1.27-.57 1.86.64 2.25 1.22c.38.57.52.7.6.9.09.2.43.19.32 0\"></path> <path fill=\"#658d5c\" d=\"M14.19 18.86c.4-.62.74-1.33 1.7-1.76.95-.43 1.1-.55 1.38-1 .28-.43.82-.73.82-.73s-.01-.41.17-.68.16-.6.16-.6.23.11.17-.3c-.05-.43.22-.7.22-.7s-.34-.45-.05-1.02c-.34-.56-.18-1.19-.18-1.19s-.3-.52 0-1.08c.93.37.98 1.57.98 1.57l.53-.74s.44.28.33.8c-.11.53-.08.64-.08.64l.5-.4s.35.57-.14 1.02c-.5.46-.59.6-.59.6s.25-.14.25.07-.05.31-.05.31l.65-.52s.2.71-.26 1.11-.74.55-.74.55.87-.24.86-.18-.18.5-.62.74a4 4 0 0 1-.78.34l.9.16s-.22.62-.75.62c-.52 0-1.02-.16-1.02-.16s.37.3.48.52c.12.23.43.25.43.25s-.29.3-.83.2-.6-.8-1.27-.57-1.86.64-2.25 1.22c-.38.57-.52.7-.6.9-.09.2-.43.19-.32 0\"></path> <path fill=\"#e1e8ed\" d=\"M12.95 17.66c-.65.07-.82.49-1.73.1-.04.2.29.5.29.5s-.5.02-.74-.06c.13.5.71.68 1.2.64.48-.04.78-.38 1.35-.26.56.12.96.28.85-.69-.38-.39-1.22-.23-1.22-.23m4.2 0c.66.07.82.49 1.74.1.04.2-.3.5-.3.5s.5.02.74-.06c-.12.5-.7.68-1.19.64-.48-.04-.79-.38-1.35-.26s-.96.28-.85-.69c.37-.39 1.22-.23 1.22-.23\"></path> <path fill=\"#f5f8fa\" d=\"M17.2 17.03c-.37-.16-1.2.24-2.15.24s-1.77-.4-2.14-.24-.3.22-.28.35l.1.7c.02.2.25.33.5.16.27-.17.6-.25 1.06-.12s.74.12.74.12.33.01.78-.12c.46-.13.8-.05 1.06.12s.48.04.51-.16l.1-.7c.02-.13.1-.18-.28-.35\"></path> <path fill=\"#99aab5\" d=\"M16.68 17.38c0 .08-.17.41-1.73.41s-1.68-.33-1.68-.4c0-.09.26.17 1.82.17s1.6-.26 1.6-.18\"></path>")
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
