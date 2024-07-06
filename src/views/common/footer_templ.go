// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package common

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Footer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"border-gray-400 dark:border-white\"><footer class=\"text-center p-4 dark:text-white dark:bg-slate-800\"><div class=\"flex justify-around items-center flex-wrap\"><div class=\"mt-4 sm:mt-0 sm:mr-4\">&copy; 2024 Antonio B. All rights reserved.</div><div class=\"flex flex-col items-center\"><a href=\"/\" class=\"mb-2 sm:mb-0\">Home</a> <a href=\"/articles\" class=\"mb-2 sm:mb-0\">Articles</a> <a href=\"/about\" class=\"mb-2 sm:mb-0\">About Me</a> <a href=\"contact\" class=\"mb-2 sm:mb-0\">Contact Me</a></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
