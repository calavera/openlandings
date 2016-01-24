function configureSite(template) {
	var md = window.markdownit();
	var result = md.render($("#markdown-content").text());
	
	$("#landing").attr("value", result);
	$("#template").attr("value", template);
	$("form").submit();
}
