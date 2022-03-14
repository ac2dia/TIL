# @NotNull vs @Column(nullable = false)

``` java
@Entity
public class Book {
 
    @Id
    @GeneratedValue
    private Long id;
 
    @Column(nullable = false)
    private String titleNullable;
 
    @NotNull
    private String titleNotNull;
 
    private LocalDate publishingDate;
     
    ... 
}
```

``` sql
CREATE TABLE public.book
(
    id bigint NOT NULL,
    publishingdate date,
    titlenotnull character varying(255) COLLATE pg_catalog."default" NOT NULL,
    titlenullable character varying(255) COLLATE pg_catalog."default" NOT NULL,
    version integer NOT NULL,
    CONSTRAINT book_pkey PRIMARY KEY (id)
)
```

- JPA 사용시 null 체크를 위해 @NotNull 및 @Column 애노테이션을 제공합니다.

- @NotNull 의 경우 제약 조건 설정 및 유효성 검사를 해준다.
- @Column 의 경우 유효성 검사는 하지 않는다.

- @Column 애노테이션을 사용하는 경우 db table과 @Entity 의 제약 조건 및 유효성 설정이 동일하지 않으면 이슈가 발생할 수 있기 때문에, @NotNull 을 통해 제약 조건과 유효성 검사를 동시에 하도록 합니다.

# Reference
[1] Hibernate Tips: What’s the difference between @Column(nullable = false) and @NotNull, https://thorben-janssen.com/hibernate-tips-whats-the-difference-between-column-nullable-false-and-notnull/
[2] [JPA] nullable = false와 @NotNull의 차이점, https://bottom-to-top.tistory.com/14