/* Code generated by cmd/cgo; DO NOT EDIT. */

#include <stdlib.h>
#include "_cgo_export.h"

#pragma GCC diagnostic ignored "-Wunknown-pragmas"
#pragma GCC diagnostic ignored "-Wpragmas"
#pragma GCC diagnostic ignored "-Waddress-of-packed-member"
#pragma GCC diagnostic ignored "-Wunknown-warning-option"
#pragma GCC diagnostic ignored "-Wunaligned-access"
extern void crosscall2(void (*fn)(void *), void *, int, size_t);
extern size_t _cgo_wait_runtime_init_done(void);
extern void _cgo_release_context(size_t);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()


#define _cgo_msan_write(addr, sz)

extern void _cgoexp_d7c74c9ef283_doubler(void *);

CGO_NO_SANITIZE_THREAD
__declspec(dllexport) GoInt doubler(GoInt i)
{
	size_t _cgo_ctxt = _cgo_wait_runtime_init_done();
	typedef struct {
		GoInt p0;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) _cgo_argtype;
	static _cgo_argtype _cgo_zero;
	_cgo_argtype _cgo_a = _cgo_zero;
	_cgo_a.p0 = i;
	_cgo_tsan_release();
	crosscall2(_cgoexp_d7c74c9ef283_doubler, &_cgo_a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return _cgo_a.r0;
}
