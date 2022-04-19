# Jackson with ObjectMapper 학습

## 학습하게 될 주요 개념

- Jackson ObjectMapper Class

- Serialization ( Java Object to JSON )
- DeSerialization ( JSON to Java Object )

## ObjectMapper 기능

### Java Object to JSON

```java

Car car = new Car("현대", "그랜저", 2022);
String carAsString = objectMapper.writeValueAsString(car);

{"brand":"현대", "model":"그랜저", "year":2022}
```

- writeValueAsString(value)
  - value: String 타입으로 변환할 대상

### JSON to Java Object

```java
String json = "{\"brand\":\"현대\", \"model\":\"그랜저\", \"year\":2022}";
Car car = objectMapper.readValue(json, Car.class);
```

- readValue(arg, type)
  - arg: 지정된 타입으로 변환할 대상
  - type: 대상을 어떤 타입으로 변환할 것인지 클래스를 명시 (Class 객체, TypeReference 객체 ...)

### JSON to Jackson JsonNode

```java
String json = "{\"brand\":\"현대\", \"model\":\"그랜저\", \"year\":2022}";

JsonNode jsonNode = objectMapper.readTree(json);
String brand = jsonNode.get("brand").asText();
String model = jsonNode.get("model").asText();
int year = jsonNode.get("year").asInt();
```

### JSON to Java Collection

```java
String json = "{\"brand\":\"현대\", \"model\":\"그랜저\", \"year\":2022}";
Map<String, Object> map = objectMapper.readValue(json, new TypeReference<Map<String, Object>>(){});
```

## 결론

- Jackson은 Java용 견고하고 성숙한 JSON 직렬화/역직렬화 라이브러리입니다.
- ObjectMapper의 API는 많은 유연성과 JSON 응답 객체를 구문 분석하고 생성하는 간단한 방법을 제공합니다.

# Reference

[1] Jackson ObjectMapper 소개, https://recordsoflife.tistory.com/599
