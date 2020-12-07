const copy = () => {
  let copyText = document.execCommand("selectAll");
  copyText.setSelectionRange(0, 99999)
  document.execCommand("copy");
  alert("Copied the team!");
}
