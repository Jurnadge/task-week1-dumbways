class testimonial {
    #quote = "";
    #image = "";
    #author = "";

    constructor(quote, image, author) {
        this.#quote = quote;
        this.#image = image;
        this.#author = author;
    }

    get quote() {
        return this.#quote;
    }

    get image() {
        return this.#image;
    }

    get author() {
        return this.#author;
    }

    get testimonialHTML() {
        return `<div class="the-actually-testimonial">
        <img src="${this.image}" />
        <p class="quote">
          ${this.quote}
        </p>
        <p class="author">~${this.author}</p>
      </div>`
    };
}

    const testimonial1 = new testimonial(
        "keren cuy",
        "https://1fid.com/wp-content/uploads/2022/12/meme-profile-picture-1024x1022.jpg",
        "Chessarjuna Mariesto"
    );

    const testimonial2 = new testimonial(
        "aku ganteng cuy",
        "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQUFlqFfHWppaEYHpo7aJgxEkd75_HoglTdHOQZnoNxRANLHAGTm8BLvp4bngbg2rheHA0&usqp=CAU",
        "Sang Pekerja Keras"
    );

    const testimonial3 = new testimonial(
        "hay maniez!!!",
        "https://i.pinimg.com/originals/2a/28/5a/2a285af8af62a57709571a27f88dabe7.jpg",
        "Xie Phalings Gantengs"
    );

    let testimonialData = [testimonial1, testimonial2, testimonial3]
    let testimonialHTML = "";

    for (let i = 0; i < testimonialData.length; i++) {
        testimonialHTML += testimonialData[i].testimonialHTML;
}
      
    document.getElementById("testimonials").innerHTML = testimonialHTML;