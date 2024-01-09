package main

import (
	"fmt"
	"strings"
)

func main() {
	response := `
$("link.favicon").attr("href", $.url('/web.shell.common/productBranding?fileName=favicon&fallbackInfo=com.informatica.tools.web.ldm.productBranding:/res/images/favicon.ico'));

var pageView = new infaw.isp.SymphonyNextLoginPageView({
    logo: "/web.shell.common/productBranding?fileName=loginPageLogo",
    logo2: "",
    namespaces: [{"id": "Native", "label": "Native"}],
    // ... (other fields)
    infaToken: "YxWbnZ-OKFtWqQYdwSm6zX346y58DBU9aNzPX6uVdCY*"
});
`
	// Find the start and end indices of infaToken
	startIndex := strings.Index(response, "infaToken:")
	if startIndex == -1 {
		fmt.Println("infaToken not found")
		return
	}
	startIndex = strings.Index(response[startIndex:], "\"") + startIndex
	endIndex := strings.Index(response[startIndex+1:], "\"") + startIndex + 1

	if endIndex == -1 {
		fmt.Println("Invalid infaToken format")
		return
	}

	infaToken := response[startIndex+1 : endIndex]
	fmt.Println("Extracted infaToken:", infaToken)
}
