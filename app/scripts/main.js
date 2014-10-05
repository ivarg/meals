(function() {
  'use strict';
  /*global $, console*/

  function openDrawer() {
    if ($('.drawer').hasClass('drawer--open')) {
      return;
    }
    toggleDrawer();
  }

  function toggleDrawer() {
    $('.drawer')
      .toggleClass('drawer--open');
  }

  function newCard(mealplan, dayNr) {
    var inputCard = $('<div></div>')
          .append('<input type="text" autofocus="false">')
          .attr('day', dayNr)
          .addClass('meal-card');
    mealplan.append(inputCard);
    inputCard.on('keypress', 'input', function(e) {
      if (e.keyCode === 13) {
        if (this.value === '') {
          inputCard.remove();
          return;
        }
        var name = $('<div></div>').append(this.value).addClass('meal-card__name');
        $(this).remove();
        inputCard.append(name);
        addCardDragHandler(0, inputCard);
      }
    });
  }

  function lift($elem) {
    console.log('lift');
    $elem.addClass('lift');
  }

  function drop($elem) {
    console.log('drop');
    $elem.removeClass('lift');
  }

  function getCurrentBase($elem) {
    var dayNr = $elem.attr('day');
    var base = null;
    $bases.forEach(function(b) {
      if (b.attr('day') === dayNr) { base = b; }
    });
    return base;
  }

  function getBaseAtDay(day) {
    return $('.meal-card__base[day="'+day+'"]');
  }

  function getCardAtDay(day) {
    return $('.meal-card[day="'+day+'"]');
  }

  function resetToDay($elem, day) {
    var base = getBaseAtDay(day);
    $elem.offset(base.offset());
    $elem.attr('day', day);
  }

  function checkBase(top, height) {
    $bases.forEach(function(b) {
      var btop = b.position().top;
      if (b !== $dragging.base &&
          top + height/2 > btop &&
          top < btop + b.height()/2) {
        $dragging.base.removeClass('dropzone');
        var srcDay = $dragging.attr('day');
        var dstDay = b.attr('day');
        $dragging.base = b;
        $dragging.base.addClass('dropzone');
        var card = getCardAtDay(dstDay);
        if (card.length === 1) {
          resetToDay(card, srcDay);
          $dragging.attr('day', dstDay);
        }
        return;
      }
    });
  }

  function addCardDragHandler($card) {
    $card.on('touchstart', function(e) {
      console.log('touchstart');
      if (e.currentTarget.classList.contains('meal-card')) {
        e.preventDefault();
        var touches = e.originalEvent.changedTouches;
        $touchId = touches[0].identifier;
        $dragging = $(touches[0].target);
        $dragging.base = getCurrentBase($dragging);
        lift($dragging);
        $dragging.removeClass('meal-card__trans');
        $ptrY = touches[0].pageY - touches[0].target.offsetTop;
      }
    });

    $card.on('touchend', function(e) {
      console.log('touchend');
      if ($dragging !== null) {
        e.preventDefault();
        drop($dragging);
        $dragging.addClass('meal-card__trans');
        resetToDay($dragging, $dragging.base.attr('day'));
        $dragging.base.removeClass('dropzone');
        $dragging = null;
        $ptrY = 0;
      }
    });


    $card.on('mousedown', function(e) {
      if (e.currentTarget.classList.contains('meal-card')) {
        $dragging = $(e.currentTarget);
        $dragging.base = getCurrentBase($dragging);
        lift($dragging);
        $dragging.removeClass('meal-card__trans');
        $ptrY = e.offsetY;
      }
    });

    $card.on('mouseup', function (e) {
      if ($dragging !== null) {
        drop($dragging);
        $dragging.addClass('meal-card__trans');
        resetToDay($dragging, $dragging.base.attr('day'));
        $dragging.base.removeClass('dropzone');
        $dragging = null;
        $ptrY = 0;
      }
    });

  }

  function layoutCard($card) {
    var d = $card.attr('day');
    var $base = getBaseAtDay(d);
    $card.offset($base.offset());
    $card.height($base.height());
    $card.width($base.width());
  }


  // dragging starts here
  var $dragging = null;
  var $bases = [];
  var $ptrY = 0;
  var $touchId = -1;

  //////////////////////
  // Document ready: start setting up stuff
  $(document).ready(function() {

    $('#whatis').mousedown(function() {
      openDrawer();
    });

    $('#config').mousedown(function() {
      openDrawer();
    });

    $('.drawer__side > h1').mousedown(function() {
      toggleDrawer();
    });

    //$('.meal-card__base').mousedown(function() {
      //newCard($(this).parent(), $(this).attr('day'));
    //});

    $.each($('.meal-card__base'), function(i, base) {
      $bases.push($(base));
    });

    $.each($('.meal-card'), function(i, card) {
      var $c = $(card);
      $c.addClass('meal-card__trans');
      layoutCard($c);
      addCardDragHandler($c);
    });

    $('.reload').on('click', function(e)  {
        console.log('reload');
    });

    $('.accept').on('click', function(e) {
        console.log('accept');
    });

    $(document.body).on("mousemove", function(e) {
      if ($dragging) {
        $dragging.offset({ top: (e.pageY-$ptrY) });
        checkBase($dragging.position().top, $dragging.height());
      }
    });

    $(document.body).on('touchmove', function(e) {
      console.log('touchmove');
      if ($dragging) {
        e.preventDefault();
        var touches = e.originalEvent.changedTouches;
        for (var i=0; i<touches.length; i++) {
          if (touches[i].identifier === $touchId) {
            $dragging.offset({ top: (touches[i].pageY-$ptrY) });
            checkBase($dragging.position().top, $dragging.height());
          }
        }
      }
    });

  });

}());

