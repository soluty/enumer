# enumer

一直在思考go语言中枚举的一个良好的使用方式，这个工具是根据我自己的枚举使用规范来生成枚举代码。

`enumer -e {enumName} -o {output} -p {package}`

该命令会生成一个int类型的type，名字为枚举名，并且默认生成了两个不需要的类型 _name_Unknown和_name_UnknownEnd

并且会有一个Valid方法可以判断枚举是否合法，并且还在枚举类型上注释添加了go:generate，可以通过stringer来生成枚举的String方法。

要新加枚举值只需要在_name_Unknown和_name_UnknownEnd中间插入值即可，然后generate一下更新string方法，所有使用到
IsValid方法的地方都不用修改。