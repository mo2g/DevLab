
<?php

//Streamline html : clear line breaks, clear tabs, remove comment tags
function streamline_html($html) {
	$arrData = preg_split( '/(<pre.*?\/pre>)/ms', $html, -1, PREG_SPLIT_DELIM_CAPTURE );
	$html = '';
	foreach ( $arrData as $str ) {
		if ( strpos( $str, '<pre' ) !== 0 ) {
			$str = preg_replace( '#/\*.+?\*/#s','', $str );//Filter script comments /* */
			$str = preg_replace( '#(?<!:)(?<!\\\\)(?<!\')(?<!")//(?<!\')(?<!").*\n#','', $str );//Filter script comments //
			$str = preg_replace( '#<!--[^\[<>].*[^\]!]-->#sU', '', $str );//Remove html comments <!-- -->
			$str = preg_replace( '#[\n\r\t]+#', ' ', $str );//Spaces replace carriage return or tab
			$str = preg_replace( '#>\s*<#', '><', $str );//Remove space between labels
			$str = preg_replace( '#\s{2,}#', ' ', $str );//Multiple spaces are merged into one space
		}
		$html .= $str;
	}
	return $html;
}