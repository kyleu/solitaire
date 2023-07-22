// Code generated by qtc from "WASM.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vutil/WASM.html:2
package vutil

//line views/vutil/WASM.html:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vutil/WASM.html:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vutil/WASM.html:2
func StreamWASMScript(qw422016 *qt422016.Writer) {
//line views/vutil/WASM.html:2
	qw422016.N().S(`
<script src="/assets/wasm/wasm_exec.js" defer="defer"></script>
<script>
  document.addEventListener("DOMContentLoaded", function() {
    if (!WebAssembly.instantiateStreaming) {
      WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
      };
    }

    const start = new Date().getTime();
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("/assets/wasm/solitaire.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
      wasmInit(new Date().getTime() - start);
    });
  });
</script>
`)
//line views/vutil/WASM.html:21
}

//line views/vutil/WASM.html:21
func WriteWASMScript(qq422016 qtio422016.Writer) {
//line views/vutil/WASM.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vutil/WASM.html:21
	StreamWASMScript(qw422016)
//line views/vutil/WASM.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vutil/WASM.html:21
}

//line views/vutil/WASM.html:21
func WASMScript() string {
//line views/vutil/WASM.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vutil/WASM.html:21
	WriteWASMScript(qb422016)
//line views/vutil/WASM.html:21
	qs422016 := string(qb422016.B)
//line views/vutil/WASM.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vutil/WASM.html:21
	return qs422016
//line views/vutil/WASM.html:21
}
