$( document ).ready(function() {
	if ($(window).width() < 620) {
	  $('#forkMeDiv').hide();
	}

	$('.js-is-hidden-alert').hide();

	$('.js-click-moreInfo').click(function(){
		if($(this).hasClass('js-show-alert')){
			$('.is-alert').addClass('alert-info');
			$('.js-is-hidden-alert').slideDown();
			$(this).removeClass('js-show-alert').addClass('js-hide-alert');
		}
		else{
			$('.is-alert').removeClass('alert-info');
			$('.js-is-hidden-alert').slideUp();
			$(this).removeClass('js-hide-alert').addClass('js-show-alert');
		}
	});

	$(window).resize(function() {
	    if ($(window).width() < 620) {
	       $('#forkMeDiv').fadeOut('slow');
	    } else {
	       $('#forkMeDiv').fadeIn('slow');
	    }
	});



});