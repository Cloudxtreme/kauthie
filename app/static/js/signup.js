Stripe.setPublishableKey(__env.stripePublishableKey);

var stripeResponseHandler = function(status, response) {
  var $form = $('form');
  console.log(response);

  if (response.error) {
    // Show the errors on the form
    $('.errors').html('<p>'+response.error.message+'</p>');
    window.location.href = '#top';
    $form.find('button').prop('disabled', false);
  } else {
    // token contains id, last4, and card type
    var token = response.id;
    // Insert the token into the form so it gets submitted to the server
    $form.append($('<input type="hidden" name="stripeToken" />').val(token));
    // and re-submit
    $form.get(0).submit();
  }
};

$(function() {
  $('input[name="cc-number"]').payment('formatCardNumber');
  $('input[name="cc-cvc"]').payment('formatCardCVC');
  $('input[name="cc-month"]').payment('restrictNumeric');
  $('input[name="cc-year"]').payment('restrictNumeric');

  $('form').submit(function(e) {
    e.preventDefault();
    $('.errors').empty();
    var $form = $(this);

    $form.find('button').prop('disabled', true);

    var errors = [];
    var validCC = $.payment.validateCardNumber($form.find('input[name="cc-number"]').val());
    var validExp = $.payment.validateCardExpiry(
        $form.find('input[name="cc-month"]').val(),
        $form.find('input[name="cc-year"]').val()
      );
    if (!validCC) {
      errors.push('Credit card number is invalid.');
    }
    if (!validExp) {
      errors.push('Credit card expiry date in invalid.');
    }

    if (errors.length > 0) {
      $form.find('button').prop('disabled', false);

      for (var i in errors) {
        $('.errors').append('<p>'+errors[i]+'</p>');
        window.location.href = '#top';
      }

      return false;
    }

    Stripe.card.createToken($form, stripeResponseHandler);
    return false;
  });
});

