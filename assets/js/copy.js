const copy = () => {
  let copyText = document.getElementById("team-pasteable");
  copyText.select();
  copyText.setSelectionRange(0, 99999)
  document.execCommand("copy");
  alert("Copied the team!");
}
