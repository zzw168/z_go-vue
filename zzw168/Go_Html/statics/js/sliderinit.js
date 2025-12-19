/*
	Author       :	Themesbazer
	Template Name:	Slider js
	Version      :	1.0
*/

(function ($) {
	"use strict";
		
		
		
		//**===================START HOME SLIDER JS===================**//	
			$('.atf-slick-slider-1').slick({
				arrows: true,
				autoplay:false,
				dots: false,
				infinite: true,
				speed: 1000,
				loop: true,
				effect: 'fade',
				slidesToShow: 1,
				adaptiveHeight: true,
				slidesToScroll: 1,
				prevArrow: '<a class="slick-prev"><i class="fa-solid fa-arrow-left-long" alt="Arrow Icon"></i></a>',
				nextArrow: '<a class="slick-next"><i class="fa-solid fa-arrow-right-long" alt="Arrow Icon"></i></a>',
			});
			
		 /* --------------------------------------------------------
			  Start CHOOSE Design
			--------------------------------------------------------- */
			$('.atf__product-slider-active').slick({
				arrows: true,
				dots: false,
				infinite: true,
				speed: 300,
				slidesToShow: 4,
				slidesToScroll: 1,
				prevArrow: '<a class="slick-prev"><i class="fa-solid fa-arrow-left-long" alt="Arrow Icon"></i></a>',
				nextArrow: '<a class="slick-next"><i class="fa-solid fa-arrow-right-long" alt="Arrow Icon"></i></a>',
				responsive: [
					{
						breakpoint: 1200,
						settings: {
							slidesToShow: 4,
							slidesToScroll: 1
						}
					},
					{
						breakpoint: 992,
						settings: {
							slidesToShow: 2,
							slidesToScroll: 1
						}
					},
					{
						breakpoint: 768,
						settings: {
							slidesToShow: 1,
							slidesToScroll: 1,
							arrows: false,
							dots: true,
						}
					},
					{
						breakpoint: 580,
						settings: {
							arrows: false,
							dots: true,
							slidesToShow: 1,
							slidesToScroll: 1
						}
					}
				]
			});
			
			
			
			 /* --------------------------------------------------------
			  Start HOT DEALS  Design
			--------------------------------------------------------- */
			$('.atf__hot-slider-active').slick({
				arrows: false,
				dots: true,
				// autoplay:true,
				infinite: true,
				speed: 300,
				slidesToShow:2,
				effect: 'fade',
				slidesToScroll: 1,
				prevArrow: '<a class="slick-prev"><i class="fa-solid fa-arrow-left-long" alt="Arrow Icon"></i></a>',
				nextArrow: '<a class="slick-next"><i class="fa-solid fa-arrow-right-long" alt="Arrow Icon"></i></a>',
				responsive: [
					{
						breakpoint: 1200,
						settings: {
							slidesToShow: 2,
							slidesToScroll: 2,
							dots: true,
							arrows: false,
						}
					},
					{
						breakpoint: 992,
						settings: {
							slidesToShow: 1,
							slidesToScroll: 1,
							dots: true,
							arrows: false,
						}
					},
					{
						breakpoint: 768,
						settings: {
							slidesToShow: 1,
							slidesToScroll: 1,
							arrows: false,
							dots: true,
						}
					},
					{
						breakpoint: 580,
						settings: {
							arrows: false,
							dots: true,
							slidesToShow: 1,
							slidesToScroll: 1
						}
					}
				]
			});
			
			 
			
		/* --------------------------------------------------------
			Star Client
		--------------------------------------------------------- */
			$('#atf-testimonial-slider').owlCarousel({
				margin:3,
				autoplay:false,
				items: 1,
				loop:true,
				nav:true,
				navText:['<i class="fa fa-angle-left"></i>','<i class="fa fa-angle-right"></i>'],
				dots:false,
				responsive:{
					0:{
						items:1
					},
					767:{
						items:1
					},
					768:{
						items:1
					},
					992:{
						items:1
					}
				}
			})
		/* --------------------------------------------------------
			END CLIENT
		--------------------------------------------------------- */

  
})(jQuery);