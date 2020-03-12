package rtda

// Thread 定义线程
type Thread struct {
	pc    int
	stack *Stack
}

// NewThread 创建一个新线程
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// PC getter
func (thread *Thread) PC() int {
	return thread.pc
}

// SetPC setter
func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

// PushFrame 帧压栈
func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

// PopFrame 帧出栈
func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

// CurrentFrame 获取当前栈帧
func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}
