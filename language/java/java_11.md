# Java 11 new features and enhancements

## 1. String API Changes

### 1.1. String.repeat(Integer)

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      String str = "1".repeat(5);

        System.out.println(str);  //11111
    }
}
```

- 해당 문자열을 n 번 반복합니다.

### 1.2. String.isBlank()

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      "1".isBlank();  //false

        "".isBlank(); //true

        "    ".isBlank(); //true
    }
}
```

- 문자열이 비어있거나 " " 공백 문자만을 포함하는 경우 true 를 반환합니다.
- 이전에는 Apache StringUtils.java 를 사용했었습니다.

### 1.3. String.strip()

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      "   hi  ".strip();  //"hi"

       "   hi  ".stripLeading();  //"hi   "

       "   hi  ".stripTrailing(); //"   hi"
    }
}
```

- 공백을 제거하고, Leading, Trailing 메서드를 이용하여 머리 공백, 꼬리 공백을 조절하여 제거할 수 있습니다.

### 1.4. String.lines()

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      String testString = "hello\nworld\nis\nexecuted";

      List<String> lines = new ArrayList<>();

      testString.lines().forEach(line -> lines.add(line));

      assertEquals(List.of("hello", "world", "is", "executed"), lines);
    }
}
```

- 여러 줄의 텍스트를 스트림으로 처리해줍니다.

# 2. Collection.toArray(IntFunction)

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      List<String> names = new ArrayList<>();
      names.add("alex");
      names.add("brian");
      names.add("charles");

      String[] namesArr1 = names.toArray(new String[names.size()]);   //Before Java 11

      String[] namesArr2 = names.toArray(String[]::new);          //Since Java 11
    }
}
```

- collection 을 array 로 변환하는 간편한 방법이 추가되었습니다.

# 3. Files.readString() and Files.writeString()

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      //Read file as string
      URI txtFileUri = getClass().getClassLoader().getResource("helloworld.txt").toURI();

      String content = Files.readString(Path.of(txtFileUri),Charset.defaultCharset());

      //Write string to file
      Path tmpFilePath = Path.of(File.createTempFile("tempFile", ".tmp").toURI());

      Path returnedFilePath = Files.writeString(tmpFilePath,"Hello World!",
                    Charset.defaultCharset(), StandardOpenOption.WRITE);
    }
}
```

- 파일 읽기, 쓰기에 대해서 보일러 플레이트 코드를 줄임으로써 훨씬 더 쉬운 방법을 사용할 수 있게 되었습니다.
- 컴퓨터 프로그래밍에서 보일러플레이트 또는 보일러플레이트 코드라고 부르는 것은 최소한의 변경으로 여러곳에서 재사용되며, 반복적으로 비슷한 형태를 띄는 코드를 말한다.

# 4. Optional.isEmpty()

```java
public class HelloWorld
{
    public static void main(String[] args)
    {
      String currentTime = null;

      assertTrue(!Optional.ofNullable(currentTime).isPresent());  //It's negative condition
      assertTrue(Optional.ofNullable(currentTime).isEmpty());   //Write it like this

      currentTime = "12:00 PM";

      assertFalse(!Optional.ofNullable(currentTime).isPresent()); //It's negative condition
      assertFalse(Optional.ofNullable(currentTime).isEmpty());  //Write it like this
    }
}
```

- Optional 은 컨테이너 객체로써 값이 없으면 객체가 비어있는 것으로 간주하고, null 에 관한 간단한 처리 방법을 제공합니다.
- isPresent() 메서드는 value 가 존재하는 경우 true, 반대의 경우 false
- isEmpty() 메서드는 value 가 존재하는 경우 false, 반대의 경우 true
- value 가 null 인 경우에 대해 처리할 수 있는 orElse(), orElseGet(), orElseThrow() 와 같은 다양한 방법이 제공됩니다.
