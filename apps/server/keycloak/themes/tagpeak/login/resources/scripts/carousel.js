class TagpeakCarousel {
    constructor(options = {}) {
        this.currentIndex = 0;
        this.templates = [];
        this.slideAutomatically = options.slideAutomatically !== false;
        this.timerInterval = options.timerInterval || 10000;
        this.bgClassSlideIndicator = options.bgClassSlideIndicator || '!bg-white';

        this.slidesContainer = document.getElementById('carouselSlides');
        this.indicatorsContainer = document.getElementById('carouselIndicators');

        if (this.slideAutomatically) {
            this.startTimer();
        }
    }

    init(templates) {
        this.templates = templates;
        this.render();
    }

    render() {
        // Clear existing content
        this.slidesContainer.innerHTML = '';
        this.indicatorsContainer.innerHTML = '';

        // Create slides
        this.templates.forEach((template) => {
            const slide = document.createElement('div');
            slide.className = 'flex-none w-full box-border';
            slide.innerHTML = template;
            this.slidesContainer.appendChild(slide);
        });

        // Create indicators
        this.templates.forEach((_, index) => {
            const indicator = document.createElement('div');
            indicator.className = `h-1 w-10 cursor-pointer ${this.slideSelectedColor(index)}`;
            indicator.addEventListener('click', () => this.selectCard(index));
            this.indicatorsContainer.appendChild(indicator);
        });

        this.updateSlidePosition();
    }

    selectCard(index) {
        this.currentIndex = index;
        this.updateSlidePosition();
        this.updateIndicators();
    }

    startTimer() {
        setInterval(() => {
            this.nextSlide();
        }, this.timerInterval);
    }

    nextSlide() {
        this.currentIndex = (this.currentIndex + 1) % this.templates.length;
        this.updateSlidePosition();
        this.updateIndicators();
    }

    updateSlidePosition() {
        this.slidesContainer.style.transform = `translateX(-${this.currentIndex * 100}%)`;
    }

    updateIndicators() {
        const indicators = this.indicatorsContainer.querySelectorAll('div');
        indicators.forEach((indicator, index) => {
            indicator.className = `h-1 w-10 cursor-pointer ${this.slideSelectedColor(index)}`;
        });
    }

    slideSelectedColor(index) {
        return this.bgClassSlideIndicator + (index === this.currentIndex ? '' : ' opacity-20');
    }
}

// Usage example:
const carousel = new TagpeakCarousel({
    slideAutomatically: true,
    timerInterval: 10000,
    bgClassSlideIndicator: '!bg-white'
});

carousel.init([
    '<div class="bg-youngBlue flex justify-center items-start h-full flex-col ">' +
    '  <img src="https://app.tagpeak.com/assets/images/carousel/slide1.png" width="90" height="90" alt="bg_slide1"' +
    '       class="w-[90%] ml-[3.75rem]">' +
    '  <div class="bg-grandis px-2 py-0.5 rounded ml-8 mb-4">' +
    '    <span class="text-sm text-youngBlue font-bold">NEW · </span>' +
    '    <span class="text-sm text-youngBlue">Refer & Earn  </span>' +
    '  </div>' +
    '  <div class="w-80 ml-8 mb-12">' +
    '    <h2 class="text-3xl text-brand">More friends, more money</h2>' +
    '    <p class="text-[15px] text-brand">' +
    '      The more the merrier — everything is better when shared and Tagpeak is no exception. Invite your friends, grow your network and get a % on their cashback forever.' +
    '    </p>' +
    '  </div>' +
    '</div>',
    '<div class="bg-youngBlue flex justify-between items-start h-full flex-col">' +
    '  <img src="https://app.tagpeak.com/assets/images/carousel/slide2.png" width="90" height="90" alt="bg_slide2"' +
    '       class="w-[80%] ml-[6.5rem] pt-[7.5rem]">' +
    '  <div class="w-80 ml-8 pb-36">' +
    '    <h2 class="text-3xl text-brand">Cash in at any time.</h2>' +
    '    <p class="text-[15px] text-brand">' +
    '      Want to cash in on your rewards? You have the power to withdraw at any time, sending the funds directly to your bank account.' +
    '    </p>' +
    '  </div>' +
    '</div>',
    '<div class="bg-youngBlue flex justify-between items-start h-full flex-col">' +
    '  <img src="https://app.tagpeak.com/assets/images/carousel/slide3.png" width="90" height="90" alt="slide3"' +
    '       class="w-[70%] ml-auto">' +
    '  <div class="pb-24 ml-8">' +
    '    <div class="bg-grandis px-2 py-0.5 rounded mb-4 w-fit">' +
    '      <span class="text-sm text-youngBlue font-bold">NEW' +
    '        · </span>' +
    '      <span class="text-sm text-youngBlue">Refer & Earn  </span>' +
    '    </div>' +
    '    <div class="w-80">' +
    '      <h2 class="text-3xl text-brand">New partnerships added</h2>' +
    '      <p class="text-[15px] text-brand">' +
    '        We are proud to welcome the following brands to the Tagpeak family.  We are excited about the opportunities these new partnerships bring — start shopping now and watch your rewards grow!' +
    '      </p>' +
    '    </div>' +
    '  </div>' +
    '</div>'
]);
