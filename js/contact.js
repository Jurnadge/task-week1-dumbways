function sendMail() {
  let name = document.getElementById("inputName").value;
  let phone = document.getElementById("inputPhone").value;
  let subject = document.getElementById("inputSubject").value;
  let message = document.getElementById("inputMessage").value;

  let emailReceiver = "chessarjunamariesto@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya, ${name} ${message}. Silahkan kirimkan pesan saya di nomor ${phone}`;
  a.click();
}
