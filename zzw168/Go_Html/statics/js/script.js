/*
	Author       :	Themesbazer
	Template Name:	Floweer - Flower & Florist Shop HTML Template
	Version      :	1.0
	
/***************************************************
==================== JS ======================
****************************************************
01. PRELOADER JS
02. ATTACHMENT IMAGE JS
03.	STICKY HEADER JS
04. MOBILE MENU 
05. SCROLL MENU
06. BACK TO TOP
07. MAILCHAMP JS
08. WOW SCROLL
09. MARQUEE JS
10. HOT DEALS JS
11. GSAP JS
12. NICE SELECT JS
13.SHOP JS 

****************************************************/

(function ($) {
    "use strict";

    /*--------------------------------------------------------------
		01.	PRELOADER JS
		--------------------------------------------------------------*/
		$(window).on('load', function() { 
			$('.loader').fadeOut();
			$('.atf-preloader').delay(350).fadeOut('slow'); 
		}); 

    /*--------------------------------------------------------------
			PRELOADER JS
		--------------------------------------------------------------*/
	/*--------------------------------------------------------------
		02.	ATTACHMENT IMAGE JS
		--------------------------------------------------------------*/
    var bg = $(".atf_attach_bg");
    bg.css("background-image", function () {
        var attach = "url(" + $(this).data("background") + ")";
        return attach;
    });

    /*--------------------------------------------------------------
		03.	STICKY HEADER JS
		--------------------------------------------------------------*/
    $(window).on("scroll", function () {
        var scroll = $(window).scrollTop();
        if (scroll < 400) {
            $(".atf-sticky-header").removeClass("atf-sticky-active");
        } else {
            $(".atf-sticky-header").addClass("atf-sticky-active");
        }
    });

    /*--------------------------------------------------------------
		04.	MOBILE MENU 
		--------------------------------------------------------------*/

    var atfHamburger = $(".atf-mobile-menu-active > ul").clone();
    var atfHamburgerMenu = $(".atf-hamburger-menu nav");
    atfHamburgerMenu.append(atfHamburger);
    if ($(atfHamburgerMenu).find(".sub-menu").length != 0) {
        $(atfHamburgerMenu).find(".sub-menu").parent().append('<button class="atf-menu-close"><i class="fas fa-chevron-right"></i></button>');
    }

    var atfSidebarMenu = $(".atf-hamburger-menu nav > ul > li button.atf-menu-close, .atf-hamburger-menu nav > ul li.menu-item-children > a");
    $(atfSidebarMenu).on("click", function (e) {
        console.log(e);
        e.preventDefault();
        if (!$(this).parent().hasClass("active")) {
            $(this).parent().addClass("active");
            $(this).siblings(".sub-menu").slideDown();
        } else {
            $(this).siblings(".sub-menu").slideUp();
            $(this).parent().removeClass("active");
        }
    });

    // Hamburger Js
    $(".atf-hamburger-toogle").on("click", function () {
        $(".atf-hamburger").addClass("atf-hamburger-open");
        $(".atf-hamburger-overlay").addClass("atf-hamburger-overlay-open");
    });

    $(".atf-hamburger-close-toggle,.atf-hamburger-overlay").on("click", function () {
        $(".atf-hamburger").removeClass("atf-hamburger-open");
        $(".atf-hamburger-overlay").removeClass("atf-hamburger-overlay-open");
    });

    /*--------------------------------------------------------------
		05.	SCROLL MENU
		--------------------------------------------------------------*/
    function scrollPage() {
        $(".atf-onepage-menu li a").click(function () {
            $(".atf-onepage-menu li a.active").removeClass("active");
            $(this).addClass("active");

            $("html, body")
                .stop()
                .animate(
                    {
                        scrollTop: $($(this).attr("href")).offset().top - 100,
                    },
                    300
                );
            return false;
        });
    }
    scrollPage();

    /*--------------------------------------------------------------
		06.	BACK TO TOP
		--------------------------------------------------------------*/

    $(window).on("scroll", function () {
        var scrolled = $(window).scrollTop();
        if (scrolled > 400) $(".back-to-top").addClass("active");
        if (scrolled < 400) $(".back-to-top").removeClass("active");
    });

    $(".back-to-top").on("click", function () {
        $("html, body").animate(
            {
                scrollTop: "0",
            },
            50
        );
    });

    /* --------------------------------------------------------
		07. MAILCHAMP JS
		--------------------------------------------------------- */
    $("#mc-form").ajaxChimp({
        url: "https://themesfamily.us22.list-manage.com/subscribe/post?u=e056d9c3aeb53b20aff997467&amp;id=e307d7e1b8&amp;f_id=0012cee1f0",
        /* Replace Your AjaxChimp Subscription Link */
    });

    /* --------------------------------------------------------
		08.	WOW SCROLL
		--------------------------------------------------------- */
    var wow = new WOW({
        //disabled for mobile
        mobile: false,
    });

    wow.init();

    /* --------------------------------------------------------
		09.	MARQUEE JS
		--------------------------------------------------------- */

    $("#marqueeLeft").marquee();

   /* --------------------------------------------------------
		10.	HOT DEALS JS
		--------------------------------------------------------- */

    function makeTimer() {
        var endTime = new Date("September 30, 2032 17:00:00 PDT");
        var endTime = Date.parse(endTime) / 1000;
        var now = new Date();
        var now = Date.parse(now) / 1000;
        var timeLeft = endTime - now;
        var days = Math.floor(timeLeft / 86400);
        var hours = Math.floor((timeLeft - days * 86400) / 3600);
        var minutes = Math.floor((timeLeft - days * 86400 - hours * 3600) / 60);
        var seconds = Math.floor(timeLeft - days * 86400 - hours * 3600 - minutes * 60);
        if (hours < "10") {
            hours = "0" + hours;
        }
        if (minutes < "10") {
            minutes = "0" + minutes;
        }
        if (seconds < "10") {
            seconds = "0" + seconds;
        }
        $("#atf-days").html(days + "<span>Days</span>");
        $("#atf-hours").html(hours + "<span>Hours</span>");
        $("#atf-minutes").html(minutes + "<span>Minutes</span>");
        $("#atf-seconds").html(seconds + "<span>Seconds</span>");
    }

    setInterval(function () {
        makeTimer();
    }, 0);

   /* --------------------------------------------------------
		11.	GSAP JS
		--------------------------------------------------------- */

    document.addEventListener("DOMContentLoaded", function () {
        // Split Content animation
        if ($(".split-content").length > 0) {
            let st = $(".split-content");
            if (st.length == 0) return;
            gsap.registerPlugin(SplitText);
            st.each(function (index, el) {
                el.split = new SplitText(el, {
                    type: "lines,words,chars",
                    linesClass: "atf-split-line",
                });
                gsap.set(el, {
                    perspective: 400,
                });
                if ($(el).hasClass("end")) {
                    gsap.set(el.split.chars, {
                        opacity: 0,
                        x: "50",
                        ease: "Back.easeOut",
                    });
                }
                if ($(el).hasClass("start")) {
                    gsap.set(el.split.chars, {
                        opacity: 0,
                        x: "-50",
                        ease: "circ.out",
                    });
                }
                if ($(el).hasClass("up")) {
                    gsap.set(el.split.chars, {
                        opacity: 0,
                        y: "80",
                        ease: "circ.out",
                    });
                }
                if ($(el).hasClass("down")) {
                    gsap.set(el.split.chars, {
                        opacity: 0,
                        y: "-80",
                        ease: "circ.out",
                    });
                }
                el.anim = gsap.to(el.split.chars, {
                    scrollTrigger: {
                        trigger: el,
                        start: "top 90%",
                    },
                    x: "0",
                    y: "0",
                    rotateX: "0",
                    scale: 1,
                    opacity: 1,
                    duration: 0.6,
                    stagger: 0.04,
                });
            });
        }

        // Image spread js
        let revealContainers = document.querySelectorAll(".spread");
        revealContainers.forEach((container) => {
            let image = container.querySelector("img");
            let tl = gsap.timeline({
                scrollTrigger: {
                    trigger: container,
                    toggleActions: "play none none none",
                },
            });

            tl.set(container, {
                autoAlpha: 1,
            });

            if (container.classList.contains("zoom-out")) {
                // Zoom-out effect
                tl.from(image, 1.5, {
                    scale: 1.4,
                    ease: Power2.out,
                });
            } else if (container.classList.contains("start") || container.classList.contains("end")) {
                let xPercent = container.classList.contains("start") ? -100 : 100;
                tl.from(container, 1.5, {
                    xPercent,
                    ease: Power2.out,
                });
                tl.from(image, 1.5, {
                    xPercent: -xPercent,
                    scale: 1,
                    delay: -1.5,
                    ease: Power2.out,
                });
            }
        });
    });
	
	 /* --------------------------------------------------------
		12.	NICE SELECT JS
		--------------------------------------------------------- */
		$('select').niceSelect();

    /* --------------------------------------------------------
		13.	SHOP JS
		--------------------------------------------------------- */
    function atf_shoping() {
        $(".atf-cart-minus").on("click", function () {
            var $input = $(this).parent().find("input");
            var count = parseInt($input.val()) - 1;
            count = count < 1 ? 1 : count;
            $input.val(count);
            $input.change();
            return false;
        });

        $(".atf-cart-plus").on("click", function () {
            var $input = $(this).parent().find("input");
            $input.val(parseInt($input.val()) + 1);
            $input.change();
            return false;
        });
		
		$("[data-width]").each(function () {
			$(this).css("width", $(this).attr("data-width"));
		});
		
		
    }
    atf_shoping();
	
})(jQuery);
