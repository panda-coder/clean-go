package presentation

func ok(body interface{}) HttpResponse {
	return HttpResponse{StatusCode: 200, Body: body}
}

func badRequest(body interface{}) HttpResponse {
	return HttpResponse{StatusCode: 400, Body: body}
}

func internalServerError(body interface{}) HttpResponse {
	return HttpResponse{StatusCode: 500, Body: body}
}
