// when click submit on contact
function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "") {
    return alert("Nama harus di isi!");
  } else if (email == "") {
    return alert("Email harus di isi!");
  } else if (phone == "") {
    return alert("Phone number harus di isi");
  } else if (subject == "") {
    return alert("Subject harus dipilih!");
  } else if (message == "") {
    return alert("Pesan harus di isi!");
  }
  let emailReceiver = "chessarjunamariesto@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Silahkan kirimkan pesan saya di nomor ${phone}`;
  a.click();

  console.log(name);
  console.log(email);
  console.log(phone);
  console.log(subject);
  console.log(message);
}
