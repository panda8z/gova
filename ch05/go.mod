module github.com/gova-jvm

go 1.13

replace (
	github.com/gova-jvm/classfile => ./classfile
	github.com/gova-jvm/classpath => ./classpath
	github.com/gova-jvm/rtda => ./rtda
	github.com/gova-jvm/instructions/base => ./instructions/base
	github.com/gova-jvm/instructions/math => ./instructions/math
	github.com/gova-jvm/instructions/stack => ./instructions/stack
	github.com/gova-jvm/instructions/loads => ./instructions/loads
	github.com/gova-jvm/instructions/conversions => ./instructions/conversions
	github.com/gova-jvm/instructions/control => ./instructions/control
	github.com/gova-jvm/instructions/comparisons => ./instructions/comparisons
	github.com/gova-jvm/instructions/constants => ./instructions/constants
	github.com/gova-jvm/instructions/extended => ./instructions/extended
)
