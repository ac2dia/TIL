# jqGrid

## colmodel option

* 위 기본 구조에서 사용한 searchResultColModel이 이에 해당하며 그 안에는 다양한 옵션이 존재한다.

```
○ name : 해당 column 에 출력해야할 데이터의 이름이다. 서버로부터 넘어오는 데이터명을 명시
○ index : 데이터 출력과는 연관이 없으며 정렬에 사용할 때 서버로 전송 시 넘어가는 데이터명이다.
○ width : column 의 가로 사이즈
○ align  : column 내에서의 데이터 정렬 (left, right, center 등?)
○ sortable : 해당 컬럼이 정렬될수 있는지를 지정한다. 일반적으로 true 설정이며 특정 컬럼의 정렬을 허용하지 않을때 사용함.
○ hidden : 특정 컬럼의 내용을 보이지 않게 하기 위해 설정한다. (input type hidden 과 용도 동일하다 생각하면 될듯.)
○ format : intefer, number, currency, data 등의 컬럼 데이터의 format을 설정하기 위해 사용된다.
○ label : colNames 가 비어있을 때 컬럼의 제목을 정의한다. (단. colName 배열과 label 속성이 없을 경우 name으로 대체한다.)
○ edittype : 편집가능한 필드 타입을 정의하며 사용한다.  (edit 관련 속성 상세 설명)
(text, textarea, select, checkbox, password, button, image, file, custom) 타입을 정의할 수 있으며 기본값은 'text'이다
```

## jqGrid option

* jqGrid를 선언해줄때 사용하는 옵션이다

```
 ○ url : "/board/searchData.do", // ajax 처럼 데이터를 주고받을 서버 url 주소이다.
 ○ datatype : "JSON", // local, xml, xmlstring, json, jsonstring, javascript 등의 데이터 타입 기재
 ○ postData : { // ajax와 마찬가지로 넘겨줄 데이터 "key1" : value1 "key2" : value2 ... },
 ○ mtype : "POST", // POST or GET 형식 선택
 ○ colNames : searchResultColNames, // 그리드 헤더의 제목 배열 (colModel 개수와 맞아야함)
 ○ colModel : searchResultColModel,// 그리드 행 데이터 (꼭 컬럼(VO)과 매칭을 시켜줘야하며 colName 개수와 일치)
 ○ rowNum : 10,  // 보여줄 행의 개수 설정
 ○ rowList : [10,20,30], // 한 화면에 보여줄 row의 수(rowNum의 수) 를 조절 가능하도록 기능 제공
 ○ pager : "#pager", // 일반적으로 페이징 처리할 태그의 id 값을 넣어줌 (대부분 pager로 통일함)
 ○ height : 260 // 그리드 높이 설정
 ○ width : 1020, // 그리드 넓이 설정 (auto 설정가능)
 ○ autowidth : true, // width 속성과 동시 사용 불가
 ○ sortname : "bbsDate", // 처음 그리드를 불러올 때 정렬에 사용할 기준 컬럼
 ○ sortorder : "desc", // 정렬 기준
 ○ sortable : true // colmodel 에기능을 사용하려 true 설정
 ○ multiselect : true // 그리드 왼쪽부분에 셀렉트 박스가 생겨 다중선택이 가능해진다.
 ○ emptyrecode : "작성된 글이 없습니다.", // 뿌려줄 데이터가 없을 경우 보여줄 문자열 지정
 ○ rownumber : true, // 각 row의 맨 앞줄에 각 행의 번호가 자동으로 부여 되도록 설정 });
```

## jqGrid event option

* jqGrid 를 사용할 때 발생하는 이벤트 제어 옵션이다. (jqGrid option 사용처럼 jqGrid 안에 넣어서 사용하면 된다.)

```
○ loadError : function(xhr, status, error){ // 불러오기 실패했을 경우 발생하는 이벤트 
//xhr : XMLHttpRequest object , status : type of error , error : exception object }, 

○ onSelectRow : function(rowid, status, e){ //row를 선택했을 때 발생하는 이벤트 
//rowid: id of the row , status : status of the selection , e: event object }, 

○ beforeRequest : function(){ // 서버로 데이터를 요청하기 직전에 발생하는 이벤트 // }, 

○ gridComplete : function(){ // 그리드가 완전히 모든 작업을 완료한 후 발생하는 이벤트 // }, 

○ loadComplete : function(){ // 서버에 요청을 보낸 직후 바로 발생하는 이벤트 // }, 

○ onPaging : function(pgButton){ // 페이징이 변경되는 경우에 발생하는 이벤트 
//pgButton : [page button] and also page number in the page input box (press Enter) , etc }, 

○ onRightClickRow : function(rowid, iRow, iCol, e){ // 마우스 우클릭 했을경우 발생하는 이벤트 
//rowid : id of the row , iRow : index of the row , iCol : index of the cell , e : event object }, 

○ onSelectAll : function(aRowids, status){ // 체크박스를 눌러 모든 row를 선택할 경우 발생하는 이벤트 (multiselection 옵션 활성화 필수) //aRowids : array of the selected rowid's , status : header check box return true or false (true if checked, false if not checked) }, 

○ onCellSelect : function(rowid, iCol, cellcontent, e){// cell 즉 column 선택을 처리하기위해 사용되는 이벤트 (cell 식별 시 사용) //rowid : id of the row , iCol : index of the cell , cellcontent : content of the cell , e : event object },
```

## jqGrid methods

* jqGrid를 사용할때 갖가지 방식으로 사용이 가능한 내장되어 있는 메서드이며, 직접적으로 jqGrid를 jQuery로 접근하여 사용이 가능하다.

``` js
$('#mainGrid').jqGrid('setCell',...);
```

## ajax 요청 실행 순서

1. beforeRequest
2. loadBeforeSend
3. serializeGridData
4. loadError
**\- Error 발생시 5\~7 과정 생략**
5. beforeProcessing
6. gridComplete
7. loadComplete

## Reference

&#91;1&#93; &#91;JavaScript&#93; jqGrid 다양한 옵션 및 사용법 모음 \(colmodel\, option\, event\, methods\)\, [https://mine-it-record.tistory.com/290](https://mine-it-record.tistory.com/290)
[2] jqGrid Events, [http://www.trirand.com/jqgridwiki/doku.php?id=wiki:events](http://www.trirand.com/jqgridwiki/doku.php?id=wiki:events)