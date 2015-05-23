package main

const mirrorlistTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>Mirror List</title>
</head>

<body>
<div id="mirrorbrain-details">
  <h2>Mirrors for <a href="{{.Url}}">{{.fileName}}</a></h2>
<div id="mirrorbrain-fileinfo">
<h3>File information</h3>
<ul>
  <li><span class="mirrorbrain-label">Filename:</span> {{.FileName}}</li>
  <li><span class="mirrorbrain-label">Path:</span> {{.Path}}</li>
  <li><span class="mirrorbrain-label">Size:</span> {{.SizeMB}} ({{.SizeBytes}})</li>
  <li><span class="mirrorbrain-label">Last modified:</span> {{.LastModified}}</li>
  <li><span class="mirrorbrain-label"><a href="{{.Url}}.sha256">SHA-256 Hash</a>:</span> <tt>{{.Sha256}}</tt></li>
  <li><span class="mirrorbrain-label"><a href="{{.Url}}.sha1">SHA-1 Hash</a>:</span> <tt>{{.Sha1}}</tt></li>
  <li><span class="mirrorbrain-label"><a href="{{.Url}}.md5">MD5 Hash</a>:</span> <tt>{{.Md5}}</tt></li>
</ul>
<p><a href="{{.Url}}" class="mirrorbrain-btn">Download file from preferred mirror</a></p>
</div>

<div id="mirrorbrain-links">
<h3>Reliable downloads</h3>
<div class="mirrorbrain-links-grp">
<h4>Metalink</h4>
<ul>
  <li><a href="{{.Url}}.meta4">{{.Url}}.meta4</a> (IETF Metalink)</li>
  <li><a href="{{.Url}}.metalink">{{.Url}}.metalink</a> (old (v3) Metalink)</li>
</ul>
</div>
</div>

<div id="mirrorbrain-mirrors">
<h3>Mirrors</h3>
<p>List of best mirrors for IP address {{.RequestIp}}, located in {{.RequestCountry}} ({{.RequestCountryCode}}), network {{.RequestNetwork}} (autonomous system {{.RequestAS}}).</p>

<div class="mirrorbrain-mirrors-grp">
<h4>Found 1 mirror very close (within the same autonomous system (AS8422)</h4>
<ul>
  <li><a href="http://mirror.netcologne.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://mirror.netcologne.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 100)</li>
</ul>
</div>

<div class="mirrorbrain-mirrors-grp">
<h4>Found 9 mirrors which handle this country (DE)</h4>
<ul>
 <li><a href="http://ftp.kullen.rwth-aachen.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://ftp.kullen.rwth-aachen.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 31)</li>
 <li><a href="http://ftp.chaos-darmstadt.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://ftp.chaos-darmstadt.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 100)</li>
 <li><a href="http://berlin.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://berlin.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 30)</li>
 <li><a href="http://mirror.eu.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://mirror.eu.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 35)</li>
 <li><a href="http://mirror.selfnet.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://mirror.selfnet.de/CCC/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 32)</li>
 <li><a href="http://ftp.fau.de/cdn.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://ftp.fau.de/cdn.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 36)</li>
 <li><a href="http://ftp.halifax.rwth-aachen.de/ccc/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://ftp.halifax.rwth-aachen.de/ccc/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 100)</li>
 <li><a href="http://31c3.mirror.speedpartner.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://31c3.mirror.speedpartner.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 33)</li>
 <li><a href="http://koeln.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://koeln.ftp.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (de, prio 20)</li>
</ul>
</div>

<div class="mirrorbrain-mirrors-grp">
<h4>Found 1 mirror in other countries, but same continent (EU)</h4>
<ul>
  <li><a href="http://c3media.vsos.ethz.ch/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://c3media.vsos.ethz.ch/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (ch, prio 100)</li>
</ul>
</div>

<div class="mirrorbrain-mirrors-grp">
<h4>Found 1 mirror in other parts of the world</h4>
<ul>
  <li><a href="http://mirror.us.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4">http://mirror.us.oneandone.net/projects/media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4</a> (us, prio 50)</li>
</ul>
</div>

</div>
<address>Powered by <a href="http://mirrorbrain.org/">MirrorBrain</a></address>
</div><!-- mirrorbrain-details -->
</body>
</html>
`
