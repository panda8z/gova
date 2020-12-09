module github.com/panda8z/gova-jvm

go 1.13

replace (
	github.com/panda8z/gova-jvm/cmd => ./cmd
	github.com/panda8z/gova-jvm/classfile => ./classfile
	github.com/panda8z/gova-jvm/classpath => ./classpath
	github.com/panda8z/gova-jvm/rtda => ./rtda
	github.com/panda8z/gova-jvm/instructions/base => ./instructions/base
	github.com/panda8z/gova-jvm/instructions/math => ./instructions/math
	github.com/panda8z/gova-jvm/instructions/stack => ./instructions/stack
	github.com/panda8z/gova-jvm/instructions/loads => ./instructions/loads
	github.com/panda8z/gova-jvm/instructions/conversions => ./instructions/conversions
	github.com/panda8z/gova-jvm/instructions/control => ./instructions/control
	github.com/panda8z/gova-jvm/instructions/comparisons => ./instructions/comparisons
	github.com/panda8z/gova-jvm/instructions/constants => ./instructions/constants
	github.com/panda8z/gova-jvm/instructions/extended => ./instructions/extended
)
