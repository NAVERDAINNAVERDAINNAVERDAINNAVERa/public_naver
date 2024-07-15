<html>
<%
	value = Eval("a")
	Execute("response.write 1")
%>
<script>
function getCookieValue() {
 return document.cookie;
}
var cookie = getCookieValue();
var index = cookie.indexOf("user");
var user = eval(document.cookie.substring(index, index+10));
</script>
</html>
