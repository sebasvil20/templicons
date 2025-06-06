// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package flags

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/sebasvil20/templicons/i"

func NorfolkIsland(props ...i.Props) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<path fill=\"#eee\" d=\"M10 0h10v24H10z\"></path> <path fill=\"#007934\" d=\"M0 24V0h10v24z\"></path> <path fill=\"#007934\" d=\"M20 24V0h10v24z\"></path> <path fill=\"#007934\" d=\"M16.82 15.02c.48-.03 1.22.15 1.76.1.55-.04.46-.07.28-.2s-.33-.04-.42.05c-.1.08-.26.08-.18 0q.1-.14-.18-.12-.25.04-.42.1c-.17.06-.25-.05-.16-.09q.13-.06 0-.08c-.07 0-.07-.05.13-.08s.35.02.47.06q.18.04.34-.05c.08-.05.13-.12.24-.11s.17-.05.03-.13a.5.5 0 0 0-.45-.04q-.17.06-.22-.03c-.04-.05-.21-.05-.38.02-.16.07-.28.15-.37.03-.05-.06-.4.01-.6.07-.21.06-.25 0-.08-.07.23-.1.37-.03.71-.15.23-.08.56.05 1.1-.23.13-.07.25-.14-.1-.15-.19 0-.27.01-.33.09s-.14.05-.13 0c0-.06-.17-.1-.3-.02q-.22.14-.32.18t-.07-.05q0-.06-.1 0l-.23.08c-.07 0-.14-.06.05-.15s.37-.2.66-.18c.39.03.47-.12.66-.11.18 0 .28-.12.03-.13-.25 0-.32.05-.38.09s-.17.06-.15-.02q.03-.1-.11-.04a1 1 0 0 1-.42.04c-.1-.03-.2-.08.05-.1.25 0 .43-.11.64-.06.2.04.3-.05.36-.04.07 0 .16-.05.03-.14-.12-.1-.44-.1-.55-.08q-.17.02-.15-.04t-.1-.06c-.1 0-.37.06-.48.1-.15.07-.26-.02-.14-.1q.19-.1.41-.09c.12.01.17.03.14-.05-.02-.03-.03-.07.1-.05s.39.04.53-.04.03-.15-.2-.15c-.22 0-.72.07-.92.19s-.22-.01-.14-.05q.11-.06.05-.1-.09-.02.08-.08c.22-.08.47-.06.75-.17.21-.1.43-.2.5-.35s-.04-.1-.16-.05q-.19.04-.2.15c0 .08-.1.08-.1-.01q.02-.16-.22-.08c-.17.06-.59.2-.74.19-.14-.02-.2-.1-.05-.11s.36-.06.49-.1c.12-.04.2-.1.1-.15s-.02-.1.08-.1.23-.09.37-.07.29-.02.1-.12c-.06-.04-.13-.09-.04-.14.1-.04.04-.17-.1-.18a1 1 0 0 0-.35.03q-.12.04-.16-.04c-.02-.05-.26-.05-.36-.03s-.29.02-.14-.05c.14-.06.44-.1.59-.1q.24 0 .38-.06c.08-.03.1-.18-.08-.18s-.41.02-.37-.05c.05-.06-.11-.1-.28-.06-.18.03-.41.06-.27-.02s.41-.18.54-.17.22-.02.14-.11c-.1-.1-.03-.13.05-.15.09 0 .1-.1-.02-.13s-.44-.08-.59 0c-.21.1-.37.1-.3.03q.09-.08-.1-.03c-.24.06-.77.2-.92.2-.07-.01-.16 0-.16-.08s.06-.08.28-.12c.43-.06.9-.18 1.11-.16q.32.04.24-.08c-.06-.08-.22-.02-.15-.12s-.06-.1-.15-.1l-.24-.08c-.08-.07-.33-.03-.46.01-.13.05-.42.11-.53.1-.14 0-.14 0-.15-.1q-.01-.05.08-.08c.14-.06.71-.07.9-.04.35.06.36-.06.29-.12-.1-.08-.4-.04-.6-.01s-.4.06-.45.02q-.07-.06.14-.1c.16-.03.54-.01.6-.04s.12-.1-.01-.13a2 2 0 0 0-.58.04c-.14.05-.3.03-.34-.01-.03-.05-.02-.16.22-.14h.75q.12-.03.05-.11c-.04-.05-.09-.1-.04-.14q.04-.05 0-.1-.03-.04.02-.07c.03-.02.06-.07-.03-.06-.15.03-.38-.04-.5 0a2 2 0 0 1-.49.06c-.04 0-.06-.08.02-.09.13-.01.38-.05.48-.09q.15-.08.31-.08c.12 0 .2-.04.09-.09a.6.6 0 0 0-.43.01c-.1.05-.3.15-.42.13q-.15-.01-.07-.2c.02-.06.16-.06.26-.05.1.02.39-.05.5-.02.11.02.21-.02.09-.1-.13-.1-.24-.06-.36-.02-.12.05-.4.1-.47.05s-.1-.22.05-.17c.15.07.52-.05.63 0s.15-.04.05-.08c-.1-.05-.39 0-.49.01q-.23.08-.23-.04c0-.06.19-.05.24-.04.14.04.34-.12.46-.09.13.03.1-.03.04-.08-.06-.04-.3-.02-.4.03a1 1 0 0 1-.29.08q-.07.01-.07-.1c0-.1.05-.12.12-.16.12-.05.14-.13.05-.11-.14.03-.19.07-.19-.1 0-.08.01-.16.13-.18.11-.02.13-.09.1-.1q-.06-.01-.14.04c-.06.02-.1.04-.1-.1q.01-.12.1-.11l.07-.02.06-.05c.02-.03 0-.07-.08-.04s-.16.04-.16-.05l-.03-.7c-.01-.1-.09-.12-.09-.01s0 .52-.02.71q-.01.1-.11.04c-.06-.03-.24-.01-.12.04s.2.07.21.16q0 .13-.15.06c-.11-.05-.25-.03-.05.07.19.11.2.12.2.23-.01.11-.08.1-.2.05-.15-.06-.27-.01-.1.06q.29.12.29.28c-.01.1-.1.15-.22 0a.6.6 0 0 0-.5-.2c-.15.01-.07.08.1.14.19.06.48.19.55.26.06.07.03.12-.13.06-.17-.06-.53-.19-.64-.15-.12.04-.17.1.03.11s.54.06.69.1c.18.07.2.2.03.15-.18-.04-.66-.13-.8-.05s-.19.16.06.12c.24-.03.54-.03.68.02.07.03.13.03.13.13q-.02.14-.24.04c-.15-.06-.6-.17-.73-.06q-.2.14.09.09c.19-.03.48 0 .62.02.3.05.27.1.08.1-.18 0-.73.02-.86-.03-.13-.04-.32-.07-.4-.04-.07.03-.01.05.05.05q.11.02.07.09c-.04.04-.07.11.1.08s.43-.03.55.02.33.1.5.11c.17.02.17.03.17.14 0 .1-.2.04-.4-.07s-.56-.18-.74-.07-.07.13.1.1c.15-.03.6.04.76.1.17.05.27.01.26.15 0 .09-.1.04-.24 0a3 3 0 0 0-1.12-.16q-.17.01-.1.05.04.05-.07.09l-.21.08c-.05.03-.05.06.09.05.13 0 .26.01.18.06-.07.05-.17.11-.02.12.4.03.97.26 1.17.36s.25.11.25.2-.11.08-.23.03c-.36-.18-.72-.14-.9-.26s-.46-.24-.63-.14c-.18.1-.15.16 0 .14.16-.01.33.01.27.08s-.1.12.15.16c.35.04.92.24 1.12.32.19.07.26.12.1.15s-.17-.09-.42-.1a1 1 0 0 1-.48-.16.5.5 0 0 0-.37-.1c-.18.02-.45.05-.55.1s-.13.12.03.13c.16 0 .23.1.11.13-.12.04-.25.06-.3.11q-.06.08.19.08c.17 0 .43.06.66.16.22.1.89.4 1.07.42.17 0 .2.03.19.11 0 .09-.07.07-.22.03a3 3 0 0 1-.7-.26q-.21-.18-.4-.22c-.13 0-.1.03-.05.08.04.04.04.14-.04.07-.14-.11-.23-.23-.4-.23q-.29-.01-.44.04-.15.09.09.13c.2.03.41.09.51.14q.13.09-.04.1-.13 0 .03.1c.1.06.17.13.09.15-.11.02-.3-.18-.4-.17-.06.01 0 .05.04.08.06.03.06.14-.03.08-.11-.08-.27-.2-.45-.13-.18.06-.27.13-.11.19.15.06.26.11.18.2-.03.03-.06.1.15.1.13.02.28.06.35.13.06.06-.02.1-.1.06-.08-.03-.26-.05-.35.03q-.1.1.18.14c.5.05 1.02.2 1.31.22q.45.05.66.13c.16.06.07.23-.05.17q-.16-.09-.35-.14-.1-.01-.05.07-.02.1-.14.02a1 1 0 0 0-.42-.2c-.13-.02-.11 0-.09.04s.01.1-.07.03a1 1 0 0 0-.52-.19c-.1.01-.09.04-.06.08s-.04.12-.1.04c-.07-.08-.2-.2-.37-.1q-.1.06-.24.1t.06.1c.36.07.5-.02.73.06.18.07.34.09.4.1.07 0 .12.03.05.06q-.11.03 0 .14c.1.07.03.17-.05.1-.13-.15-.32-.1-.4-.17q-.1-.1-.1 0 0 .08-.1.01c-.24-.17-.53-.13-.66-.05s-.1.1.13.14c.65.09 1.45.21 2.14.48.38.14.36.17.34.22q-.03.08-.24-.04c-.18-.1-.63-.17-.83-.28s-.26-.06-.2.01c.05.08.03.15-.08.06-.12-.1-.36-.22-.5-.22-.13 0-.03.05.04.1.07.04.08.16-.03.1-.12-.07-.47-.25-.74-.19-.27.07-.23.1-.1.18.14.09.27.15.57.14s.58.04.72.12c.13.08 0 .1-.08.05a1 1 0 0 0-.43-.07c-.1 0-.06.05-.02.1.04.04 0 .17-.08.09s-.3-.18-.44-.21c-.29-.06-.42.06-.29.1q.13.04.07.12c-.05.06-.07.12.09.1a4 4 0 0 1 1.42.07c.32.09.25.15.16.15s-.17-.05-.13.02.04.14-.07.08q-.3-.18-.67-.2c-.08 0-.04.05 0 .09.02.04-.02.1-.1.04a1 1 0 0 0-.35-.14q-.14 0-.05.07c.05.04.04.11-.06.05-.14-.1-.37-.2-.53-.18-.17.01-.1.08.06.14.28.1.58.18.79.2.2.01.4.06.48.07.08 0 .12.05 0 .06-.16.02-.42-.09-.55-.03-.06.03.02.07.03.09.02.04.01.14-.07.08-.25-.19-.41-.16-.5-.22s-.18-.03-.12.04-.07.05-.16 0a.6.6 0 0 0-.36-.07c-.1.03-.15.06.06.17q.28.15.6.21c.2.04.23.15.12.14-.12-.02-.14.05-.06.1q.06.06.05.14-.01.07-.17-.06c-.22-.23-.71-.24-.96-.18-.24.07-.14.17.04.33.18.15.34.31.8.26.45-.05 1.26-.21 1.64-.12s.33.24.21.26c-.2.03-.2-.15-.4-.16-.22-.01-.28.08-.35.15s-.12.13-.34.03c-.27-.12-.48.05-.6.03.12.06.3.19.25.24s-.13-.05-.3-.1a.8.8 0 0 0-.6.02c-.18.1.01.2.22.25s.65.05 1.01-.04c.37-.1.82-.16.98-.16q.2.01.1.13c.13.03.28 0 .36-.05.2-.11.17 0 .13.08.13 0 .36-.15.46-.09.02.51-.01 1.55-.01 1.9 0 .06.05.05.39.05.4 0 .43.03.42-.15s-.13-1.23-.1-1.62c.04-.38.16-.35.3-.32.16.02.5.13.68.12-.1-.1-.12-.18-.02-.16a.4.4 0 0 0 .34-.09c-.08 0-.14-.1 0-.12s.2.04.44.12c.25.07 1.18.24 1.71-.07.16-.09.03-.13-.16-.13q-.28.01-.38.07-.17.07-.16 0c.02-.07-.15-.1-.3-.04-.15.07-.27.1-.24.02q.07-.1-.22-.08c-.17.01-.68-.04-.66-.12q.01-.07.28-.03c.4.08 1.3.1 1.58.09.45-.03.32-.16.15-.2s-.37-.01-.4.03c-.05.05-.2.08-.16 0 .05-.1-.2-.13-.42-.02-.14.08-.17.04-.12-.02q.06-.1-.2-.09c-.24.02-.67.02-.84-.04q-.29-.1.17-.12m-1.47-5.89q-.1.05-.12-.02c0-.06.07-.04.12-.03.05 0 .1.03 0 .05m0-.26q.1.03.2 0c.07 0 .15.04.05.07-.18.04-.37.08-.37-.03 0-.06.07-.06.12-.04m-.53-.05c-.1-.04-.07-.08-.02-.1s.15-.01.15.09c0 .06-.09.02-.13 0m-1.68 6.36c-.05-.07.01-.03.12 0q.12.04.13.1c-.01.09-.2-.03-.25-.1m1.24-1.14c.1.06.44.05.44.18 0 .05-.07.02-.15.02a1 1 0 0 1-.37-.15c-.1-.1-.02-.12.08-.05m.42 1.22c-.02.11-.22.03-.38-.05-.15-.08.01-.11.08-.1q.11.03.22 0c.08-.01.1 0 .08.15m-.1-.33c-.11 0-.35-.09-.56-.17-.21-.07-.02-.12.1-.07s.34.05.47.05c.14 0 .16.19-.01.19m-.18-2.6c-.18-.02-.34-.15-.83-.25-.2-.04-.15-.07-.01-.07s.55.17 1 .24c.1.01.21.03.2.15 0 .13-.19-.05-.36-.07m.18-.46a3 3 0 0 0-.62-.21c-.14-.03-.22-.07-.09-.09q.23-.02.48.1c.1.04.25.1.31.1.07 0 .12 0 .12.12 0 .13-.09.05-.2-.02m-.64-2.5c.18.07.54.05.78.18.1.05-.04.07-.18.01-.15-.05-.31-.02-.6-.1-.24-.08-.14-.13 0-.08m.47.96q-.31-.11-.28-.16c.01-.03.21.02.35.06.25.07.3.08.29.18 0 .08-.14 0-.36-.08m-.5.24c.21 0 .48.11.66.2.13.05.2.07.2.14 0 .06-.08.03-.15 0-.06-.04-.5-.16-.7-.24q-.34-.11-.01-.1m.88.64c0 .05-.06.01-.1.01-.1 0-.32-.15-.62-.2-.15-.02-.33-.13-.5-.18s-.03-.12.06-.07q.06.04.3.07c.16.03.52.16.63.23.05.03.24.02.23.14m.03-1.36q-.01.14-.12.05-.1-.07-.35-.12c-.15-.04-.09-.13.03-.08q.18.08.4.06c.07 0 .04.04.04.1m.27.3c0-.14.02-.14.11-.12.1.02.41 0 .54 0 .12-.02.22.03.01.06q-.22.04-.46.11c-.09.04-.2.07-.2-.06m1.23 2.65c-.09.08-.55.16-.82.28-.23.1-.28.06-.28-.07 0-.14.16-.05.24-.1s.57-.1.72-.17.22-.03.14.06m-.38.69c-.15.02-.28.09-.52.16s-.2-.04-.1-.1c.1-.05.42-.12.6-.13s.17.05.02.07m-.1-.87a1 1 0 0 1-.35.13q-.27.07-.27-.06c0-.1.06-.07.19-.06s.38.02.44 0m-.42-.5c-.07.04-.22.06-.21 0q0-.07.03-.06c.19-.01.34-.1.44-.11.15-.02.23-.03.15.03s-.34.1-.41.14m-.13-1.29c.07 0 .28-.03.37-.07q.14-.03.04.03l-.33.12q-.2.07-.2-.02c0-.07.02-.06.12-.06m.02.51c.13 0 .39 0 .47-.04l.24-.04q.07 0-.05.08c-.08.05-.36.05-.46.1a1 1 0 0 1-.28.07q-.08.02-.07-.06c0-.1.02-.11.15-.1m-.14.43q-.02-.1.07-.08c.17 0 .47-.13.58-.13.12 0 .17.06.04.07a2 2 0 0 0-.5.13c-.07.03-.2.07-.2 0m.23.54c.17.03.2.08-.03.12-.12.01-.16.02-.17-.07 0-.06.03-.07.2-.05m-.08 1c.09 0 .28-.03.34-.05.12-.03.1.01.03.05a1 1 0 0 1-.4.09c-.14-.01-.1-.1.03-.1m.29 1.97a.4.4 0 0 0-.22 0q-.13 0-.14-.13c.08.06.28-.03.46.06.08.03-.02.1-.1.07m.08-.44c-.1 0-.23.09-.36.07-.12-.03-.1-.13.02-.1.1.02.17-.02.3-.05.14-.03.13.08.04.08m-.37-.43.2-.02q.11-.01.01.05c-.05.03-.13.08-.21.06-.09-.02-.07-.09 0-.09m.05-.48c-.24.1-.16-.05-.05-.1l.59-.2c.2-.06.33-.06.48-.08q.2-.04.04.11c-.1.11-.21.08-.37.07s-.45.1-.69.2m.85.17q-.21-.02-.06-.06c.06-.02.17-.04.21-.01q.08.06-.15.07\"></path>")
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
