use strict;
use warnings;

use Tk;
my $top = MainWindow->new;
my $label = $top->Label(
    -text => 'Hello World!',
    -font => ['fixed', 20]
);
$label->pack;
MainLoop;
