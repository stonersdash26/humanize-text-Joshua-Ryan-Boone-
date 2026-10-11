func verifyToken(token, ip string) (bool, []string, error) {
  form := url.Values{
    "secret":   {"your_secret_key"},
    "response": {token},
    "remoteip": {ip},
    "sitekey":  {"d2a034a6-03e6-4ef2-91d4-39dc9fbc5b1c"},
  }
  resp, err := http.PostForm(
    "https://api.hcaptcha.com/siteverify",
    form,
  )
  if err != nil { return false, nil, err }
  defer resp.Body.Close()

  var out struct {
    Success bool `json:"success"`
    ErrorCodes []string `json:"error-codes"`
  }
  if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
    return false, nil, err
  }
  if out.Success { return true, []string{}, nil }
  return false, out.ErrorCodes, nil
}