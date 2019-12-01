package mimetype

import "github.com/gabriel-vasile/mimetype/internal/matchers"

// root is a matcher which passes for any slice of bytes.
// When a matcher passes the check, the children matchers
// are tried in order to find a more accurate mime type.
var root = newNode("application/octet-stream", "", matchers.True,
	sevenZ, zip, pdf, ole, ps, psd, ogg, png, jpg, jp2, jpx, jpm, gif, webp, exe, elf,
	ar, tar, xar, bz2, fits, tiff, bmp, ico, mp3, flac, midi, ape, musePack, amr,
	wav, aiff, au, mpeg, quickTime, mqv, mp4, webM, threeGP, threeG2, avi, flv,
	mkv, asf, aac, voc, aMp4, m4a, txt, gzip, class, swf, crx, woff, woff2, otf,
	eot, wasm, shx, dbf, dcm, rar, djvu, mobi, lit, bpg, sqlite3, dwg, nes, macho,
	qcp, icns, heic, heicSeq, heif, heifSeq, mrc, mdb, accdb, zstd,
)

// The list of nodes appended to the root node
var (
	gzip = newNode("application/gzip", "gz", matchers.Gzip).
		alias("application/x-gzip", "application/x-gunzip", "application/gzipped", "application/gzip-compressed", "application/x-gzip-compressed", "gzip/document")
	sevenZ = newNode("application/x-7z-compressed", "7z", matchers.SevenZ)
	zip    = newNode("application/zip", "zip", matchers.Zip, xlsx, docx, pptx, epub, jar, odt, ods, odp, odg, odf).
		alias("application/x-zip", "application/x-zip-compressed")
	tar = newNode("application/x-tar", "tar", matchers.Tar)
	xar = newNode("application/x-xar", "xar", matchers.Xar)
	bz2 = newNode("application/x-bzip2", "bz2", matchers.Bz2)
	pdf = newNode("application/pdf", "pdf", matchers.Pdf).
		alias("application/x-pdf")
	xlsx = newNode("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "xlsx", matchers.Xlsx)
	docx = newNode("application/vnd.openxmlformats-officedocument.wordprocessingml.document", "docx", matchers.Docx)
	pptx = newNode("application/vnd.openxmlformats-officedocument.presentationml.presentation", "pptx", matchers.Pptx)
	epub = newNode("application/epub+zip", "epub", matchers.Epub)
	jar  = newNode("application/jar", "jar", matchers.Jar)
	ole  = newNode("application/x-ole-storage", "", matchers.Ole, xls, pub, ppt, doc)
	doc  = newNode("application/msword", "doc", matchers.Doc).
		alias("application/vnd.ms-word")
	ppt = newNode("application/vnd.ms-powerpoint", "ppt", matchers.Ppt).
		alias("application/mspowerpoint")
	pub = newNode("application/vnd.ms-publisher", "pub", matchers.Pub)
	xls = newNode("application/vnd.ms-excel", "xls", matchers.Xls).
		alias("application/msexcel")
	ps   = newNode("application/postscript", "ps", matchers.Ps)
	fits = newNode("application/fits", "fits", matchers.Fits)
	ogg  = newNode("application/ogg", "ogg", matchers.Ogg, oggAudio, oggVideo).
		alias("application/x-ogg")
	oggAudio = newNode("audio/ogg", "oga", matchers.OggAudio)
	oggVideo = newNode("video/ogg", "ogv", matchers.OggVideo)
	txt      = newNode("text/plain", "txt", matchers.Txt, html, svg, xml, php, js, lua, perl, python, json, ndJson, rtf, tcl, csv, tsv, vCard, iCalendar, warc)
	xml      = newNode("text/xml; charset=utf-8", "xml", matchers.Xml, rss, atom, x3d, kml, xliff, collada, gml, gpx, tcx, amf, threemf)
	json     = newNode("application/json", "json", matchers.Json, geoJson)
	csv      = newNode("text/csv", "csv", matchers.Csv)
	tsv      = newNode("text/tab-separated-values", "tsv", matchers.Tsv)
	geoJson  = newNode("application/geo+json", "geojson", matchers.GeoJson)
	ndJson   = newNode("application/x-ndjson", "ndjson", matchers.NdJson)
	html     = newNode("text/html; charset=utf-8", "html", matchers.Html)
	php      = newNode("text/x-php; charset=utf-8", "php", matchers.Php)
	rtf      = newNode("text/rtf", "rtf", matchers.Rtf)
	js       = newNode("application/javascript", "js", matchers.Js).
			alias("application/x-javascript", "text/javascript")
	lua    = newNode("text/x-lua", "lua", matchers.Lua)
	perl   = newNode("text/x-perl", "pl", matchers.Perl)
	python = newNode("application/x-python", "py", matchers.Python)
	tcl    = newNode("text/x-tcl", "tcl", matchers.Tcl).
		alias("application/x-tcl")
	vCard     = newNode("text/vcard", "vcf", matchers.VCard)
	iCalendar = newNode("text/calendar", "ics", matchers.ICalendar)
	svg       = newNode("image/svg+xml", "svg", matchers.Svg)
	rss       = newNode("application/rss+xml", "rss", matchers.Rss).
			alias("text/rss")
	atom    = newNode("application/atom+xml", "atom", matchers.Atom)
	x3d     = newNode("model/x3d+xml", "x3d", matchers.X3d)
	kml     = newNode("application/vnd.google-earth.kml+xml", "kml", matchers.Kml)
	xliff   = newNode("application/x-xliff+xml", "xlf", matchers.Xliff)
	collada = newNode("model/vnd.collada+xml", "dae", matchers.Collada)
	gml     = newNode("application/gml+xml", "gml", matchers.Gml)
	gpx     = newNode("application/gpx+xml", "gpx", matchers.Gpx)
	tcx     = newNode("application/vnd.garmin.tcx+xml", "tcx", matchers.Tcx)
	amf     = newNode("application/x-amf", "amf", matchers.Amf)
	threemf = newNode("application/vnd.ms-package.3dmanufacturing-3dmodel+xml", "3mf", matchers.Threemf)
	png     = newNode("image/png", "png", matchers.Png)
	jpg     = newNode("image/jpeg", "jpg", matchers.Jpg)
	jp2     = newNode("image/jp2", "jp2", matchers.Jp2)
	jpx     = newNode("image/jpx", "jpf", matchers.Jpx)
	jpm     = newNode("image/jpm", "jpm", matchers.Jpm).
		alias("video/jpm")
	bpg  = newNode("image/bpg", "bpg", matchers.Bpg)
	gif  = newNode("image/gif", "gif", matchers.Gif)
	webp = newNode("image/webp", "webp", matchers.Webp)
	tiff = newNode("image/tiff", "tiff", matchers.Tiff)
	bmp  = newNode("image/bmp", "bmp", matchers.Bmp).
		alias("image/x-bmp", "image/x-ms-bmp")
	ico  = newNode("image/x-icon", "ico", matchers.Ico)
	icns = newNode("image/x-icns", "icns", matchers.Icns)
	psd  = newNode("image/vnd.adobe.photoshop", "psd", matchers.Psd).
		alias("image/x-psd", "application/photoshop")
	heic    = newNode("image/heic", "heic", matchers.Heic)
	heicSeq = newNode("image/heic-sequence", "heic", matchers.HeicSequence)
	heif    = newNode("image/heif", "heif", matchers.Heif)
	heifSeq = newNode("image/heif-sequence", "heif", matchers.HeifSequence)
	mp3     = newNode("audio/mpeg", "mp3", matchers.Mp3).
		alias("audio/x-mpeg", "audio/mp3")
	flac = newNode("audio/flac", "flac", matchers.Flac)
	midi = newNode("audio/midi", "midi", matchers.Midi).
		alias("audio/mid", "audio/sp-midi", "audio/x-mid", "audio/x-midi")
	ape      = newNode("audio/ape", "ape", matchers.Ape)
	musePack = newNode("audio/musepack", "mpc", matchers.MusePack)
	wav      = newNode("audio/wav", "wav", matchers.Wav).
			alias("audio/x-wav", "audio/vnd.wave", "audio/wave")
	aiff = newNode("audio/aiff", "aiff", matchers.Aiff)
	au   = newNode("audio/basic", "au", matchers.Au)
	amr  = newNode("audio/amr", "amr", matchers.Amr).
		alias("audio/amr-nb")
	aac  = newNode("audio/aac", "aac", matchers.Aac)
	voc  = newNode("audio/x-unknown", "voc", matchers.Voc)
	aMp4 = newNode("audio/mp4", "mp4", matchers.AMp4).
		alias("audio/x-m4a", "audio/x-mp4a")
	m4a  = newNode("audio/x-m4a", "m4a", matchers.M4a)
	mp4  = newNode("video/mp4", "mp4", matchers.Mp4)
	webM = newNode("video/webm", "webm", matchers.WebM).
		alias("audio/webm")
	mpeg      = newNode("video/mpeg", "mpeg", matchers.Mpeg)
	quickTime = newNode("video/quicktime", "mov", matchers.QuickTime)
	mqv       = newNode("video/quicktime", "mqv", matchers.Mqv)
	threeGP   = newNode("video/3gpp", "3gp", matchers.ThreeGP).
			alias("video/3gp", "audio/3gpp")
	threeG2 = newNode("video/3gpp2", "3g2", matchers.ThreeG2).
		alias("video/3g2", "audio/3gpp2")
	avi = newNode("video/x-msvideo", "avi", matchers.Avi).
		alias("video/avi", "video/msvideo")
	flv = newNode("video/x-flv", "flv", matchers.Flv)
	mkv = newNode("video/x-matroska", "mkv", matchers.Mkv)
	asf = newNode("video/x-ms-asf", "asf", matchers.Asf).
		alias("video/asf", "video/x-ms-wmv")
	class   = newNode("application/x-java-applet; charset=binary", "class", matchers.Class)
	swf     = newNode("application/x-shockwave-flash", "swf", matchers.Swf)
	crx     = newNode("application/x-chrome-extension", "crx", matchers.Crx)
	woff    = newNode("font/woff", "woff", matchers.Woff)
	woff2   = newNode("font/woff2", "woff2", matchers.Woff2)
	otf     = newNode("font/otf", "otf", matchers.Otf)
	eot     = newNode("application/vnd.ms-fontobject", "eot", matchers.Eot)
	wasm    = newNode("application/wasm", "wasm", matchers.Wasm)
	shp     = newNode("application/octet-stream", "shp", matchers.Shp)
	shx     = newNode("application/octet-stream", "shx", matchers.Shx, shp)
	dbf     = newNode("application/x-dbf", "dbf", matchers.Dbf)
	exe     = newNode("application/vnd.microsoft.portable-executable", "exe", matchers.Exe)
	elf     = newNode("application/x-elf", "", matchers.Elf, elfObj, elfExe, elfLib, elfDump)
	elfObj  = newNode("application/x-object", "", matchers.ElfObj)
	elfExe  = newNode("application/x-executable", "", matchers.ElfExe)
	elfLib  = newNode("application/x-sharedlib", "so", matchers.ElfLib)
	elfDump = newNode("application/x-coredump", "", matchers.ElfDump)
	ar      = newNode("application/x-archive", "a", matchers.Ar, deb).
		alias("application/x-unix-archive")
	deb = newNode("application/vnd.debian.binary-package", "deb", matchers.Deb)
	dcm = newNode("application/dicom", "dcm", matchers.Dcm)
	odt = newNode("application/vnd.oasis.opendocument.text", "odt", matchers.Odt, ott).
		alias("application/x-vnd.oasis.opendocument.text")
	ott = newNode("application/vnd.oasis.opendocument.text-template", "ott", matchers.Ott).
		alias("application/x-vnd.oasis.opendocument.text-template")
	ods = newNode("application/vnd.oasis.opendocument.spreadsheet", "ods", matchers.Ods, ots).
		alias("application/x-vnd.oasis.opendocument.spreadsheet")
	ots = newNode("application/vnd.oasis.opendocument.spreadsheet-template", "ots", matchers.Ots).
		alias("application/x-vnd.oasis.opendocument.spreadsheet-template")
	odp = newNode("application/vnd.oasis.opendocument.presentation", "odp", matchers.Odp, otp).
		alias("application/x-vnd.oasis.opendocument.presentation")
	otp = newNode("application/vnd.oasis.opendocument.presentation-template", "otp", matchers.Otp).
		alias("application/x-vnd.oasis.opendocument.presentation-template")
	odg = newNode("application/vnd.oasis.opendocument.graphics", "odg", matchers.Odg, otg).
		alias("application/x-vnd.oasis.opendocument.graphics")
	otg = newNode("application/vnd.oasis.opendocument.graphics-template", "otg", matchers.Otg).
		alias("application/x-vnd.oasis.opendocument.graphics-template")
	odf = newNode("application/vnd.oasis.opendocument.formula", "odf", matchers.Odf).
		alias("application/x-vnd.oasis.opendocument.formula")
	rar = newNode("application/x-rar-compressed", "rar", matchers.Rar).
		alias("application/x-rar")
	djvu    = newNode("image/vnd.djvu", "djvu", matchers.DjVu)
	mobi    = newNode("application/x-mobipocket-ebook", "mobi", matchers.Mobi)
	lit     = newNode("application/x-ms-reader", "lit", matchers.Lit)
	sqlite3 = newNode("application/x-sqlite3", "sqlite", matchers.Sqlite)
	dwg     = newNode("image/vnd.dwg", "dwg", matchers.Dwg).
		alias("image/x-dwg", "application/acad", "application/x-acad", "application/autocad_dwg", "application/dwg", "application/x-dwg", "application/x-autocad", "drawing/dwg")
	warc  = newNode("application/warc", "warc", matchers.Warc)
	nes   = newNode("application/vnd.nintendo.snes.rom", "nes", matchers.Nes)
	macho = newNode("application/x-mach-binary", "macho", matchers.MachO)
	qcp   = newNode("audio/qcelp", "qcp", matchers.Qcp)
	mrc   = newNode("application/marc", "mrc", matchers.Marc)
	mdb   = newNode("application/x-msaccess", "mdb", matchers.MsAccessMdb)
	accdb = newNode("application/x-msaccess", "accdb", matchers.MsAccessAce)
	zstd  = newNode("application/zstd", "zst", matchers.Zstd)
)
