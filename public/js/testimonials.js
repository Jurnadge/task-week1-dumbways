const testimonialPromise = new Promise(
  (testimonialResolve, testimonialReject) => {
    const testimonialData = new XMLHttpRequest();
    testimonialData.open(
      "GET",
      "https://api.npoint.io/449ff1485066401b7a5f",
      true
    );
    testimonialData.onload = () => {
      if (testimonialData.status === 200) {
        testimonialResolve(JSON.parse(testimonialData.response));
      } else {
        testimonialReject("error loading data");
      }
    };
    testimonialData.onerror = () => {
      testimonialReject("network eror");
    };
    testimonialData.send();
  }
);

async function allTestimonials() {
  const testimonialResponse = await testimonialPromise;

  let testimonialHTML = "";
  testimonialResponse.forEach(function (item) {
    testimonialHTML += `<div
                            class="col-md-5 card me-auto ms-auto shadow-lg pb-3 mb-3"
                            style="width: 25rem;"
                        >
                            <img
                            class="object-fit-contain"
                            src="${item.image}"
                            />
                            <p class="text-start mt-2">
                            "${item.quote}"
                            </p>
                            <p class="text-end mt-4 mb-2">
                            ~${item.author}
                            </p>
                            <p class="text-end mt-2">
                            ${item.rating} <i class="fa-solid fa-star"></i>
                            </p>
                        </div>`;
  });

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

allTestimonials();

async function filterTestimonials(rating) {
  const testimonialResponse = await testimonialPromise;

  let testimonialHTML = "";
  const testimonialFiltered = testimonialResponse.filter((item) => {
    return item.rating === rating;
  });

  if (testimonialFiltered.length === 0) {
    testimonialHTML = `<h1>Data Not Found<h1>`;
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `<div
                            class="col-md-5 card me-auto ms-auto shadow-lg pb-3 mb-3"
                            style="width: 25rem;"
                            >
                                <img
                                class="object-fit-contain"
                                src="${item.image}"
                                />
                                <p class="text-start mt-2">
                                "${item.quote}"
                                </p>
                                <p class="text-end mt-4 mb-2">
                                ~${item.author}
                                </p>
                                <p class="text-end mt-2">
                                ${item.rating} <i class="fa-solid fa-star"></i>
                                </p>
                            </div>`;
    });
  }

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}
