package heraldry

import svg "github.com/ajstarks/svgo"

func renderBearHeadCoupedToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M259.1,231.4c1.5,5.2,1.3,9.7-2.1,14.2c-6.7,8.8-1.9,15.8,5.2,20.4c2,1.3,4.1,1.6,6.4,0c1.3-1,3.2-3.4,4.8-1    c1.4,2.2,1.6,5.3-1.1,7.1c-2.5,1.7-5.4,2.9-8,4.3c0.7,1.9,4.6,2.8,1.8,6.1c-3.7,4.4-7.5,2.2-11.6,0.3c0.7,3.7,0.6,7.3,4.7,8.8    c0.6,0.2,1,0.9,0.7,1.6c-0.4,1.1-1.4,0.6-2.1,0.7c-7.5,0.7-14.8-1.2-22.3-1c-3.8,0.1-7.4,0.9-10.9,2.1c-6.1,2-12.6,1.4-18.7,0.7    c-6.1-0.7-12.2-0.5-18.3-0.7c-6.2-0.2-12.3,0-18.5,0c-3.4,0-6.7,0.2-10.1,1.4c-3.3,1.2-6.9,1-10.6-0.2c-8.9-3.1-18-2.9-27.1-1.2    c-1,0.2-2-0.1-2.7,0.9c-2.8,4.1-6.9,4.1-10.4,2c-5.6-3.4-12.9,0.1-17.6-4.2c-5.6-5.2-8.9-3.5-13.6,1.2c-4.7,4.8-11.2,6-18.7,4.7    c7.3-5.6,0.3-11,0.5-16.5c0.1-1.7-1.8-0.9-2.8-1c-3-0.2-6-0.1-8.9-0.5c-0.6-0.1-2.3-0.3-2.2-2.1c0.1-1.4,1.1-1.4,1.9-1.8    c0.6-0.3,1.2-0.7,1.8-0.8c7.9-0.6,10.5-6.7,13.7-12.5c2.1-3.8,4.8-7.2,7.4-11c-3-1.2-6.6,1-9.2-2c-0.6-0.7-1.7-0.7-1.6-2.1    c0.2-1.6,1.3-1.1,2.2-1.2c9-1.2,11.4-4.9,12.9-14.1c1-6,3-11.9,3.5-18c0.4-4.7,4.3-7.1,7.2-10.3c-3.2-2-6.8-3.4-9.2-6.8    c1.7-1.2,3.3-0.9,4.6-0.7c7.7,0.9,12.1-1.1,15.4-8.4c4.6-10.3,12.5-15.3,23.6-15.6c1.6,0,2.8,0,2.8-2.3c0-2.2-0.6-2.8-2.8-2.8    c-9.8,0.1-19.7,0.3-29.5,0c-8.7-0.3-16.9-2.6-23.3-9.1c-1.9-1.9-3.2-1.3-4.8,0c-0.4,0.3-0.7,0.7-1.1,1.1c-2.7,2.7-5.7,4.9-9.3,1.8    c-2.4-2-1.6-7.4,1.1-10.7c5.2-6.1,11.4-10.5,19.6-12.1c-7.6-3.2-15.9-4-20.8-10.5c-3.3-4.3-6.8-9-7.3-14.9    c-0.3-4.7,1.9-7.2,6.2-8.6c13.2-4,26.4-7.9,39.5-12.3c18.2-6.2,36.1-13.3,52.6-23.5c8.3-5.1,17.6-8.7,27.1-11    c6.4-1.5,9.9-5.2,11.9-10.9c2.6-7.4,8.9-10.2,15.4-12c5.7-1.6,11.5-2.7,16.7-5.8c4.6-2.7,8.6-2.4,7.8,4.5c-0.6,5,0.5,10.2-1,15.3    c-0.3,1.2-0.9,2.9,1.3,4.2c1.9,1.1,3,0.1,3.9-0.5c5.6-3.3,11.4-5.1,18-4.9c5.4,0.1,10.9,0.9,16.3-0.8c2.2-0.7,4.9-0.3,6,2.2    c1,2.2,0.6,4.5-1.5,6.2c-2.3,1.9-3.1,4.5-3.1,7.4c0.1,4.2-1.1,8.2-2.6,11.8c-1.4,3.5-1,4.3,2.2,5.8c4,1.9,6.3,5.4,6.4,10    c0.1,3.8,0.2,7.7,0,11.5c-0.4,7.9,1.4,14.2,9.9,16.4c0.4,2.2-1.8,2.4-2.1,2.3c-7-0.5-5.5,4.4-4.4,7.8c0.9,2.9,0.5,5.3-0.3,8    c-2.3,7.9,0.3,14.9,4.5,21.4c1.3,2,4.3,3.8,1.7,6.2c-2.2,2-5.6,1.6-8.3,0.2c-1.2-0.6-1.7-2.4-3.1-4.4c-0.9,5.8,0.2,9.5,3,12.4    c1.3,1.3,2.6,2.6,3.9,3.9c0.8,0.8,1.7,1.6,1.3,3c-0.5,1.7-2,0.9-2.9,0.9c-3.6-0.2-5-0.3-4.3,4.8c1.2,8.7-1.6,17.4-3.1,26    c-1,5.9-0.1,11,4.7,15.1c1.2,1,2.8,1.9,2.7,3.9c-0.1,2.1,0.3,4.2-3.1,2.7C264.6,232.5,262,232.1,259.1,231.4z"
	pathData1 := "M92.1,241.1c-1.7-0.6-3,0.8-5.1,0.8c2.3-3,4.4-5.7,8.4-7.7c-4.3-0.5-6.3,1.3-8.5,3.8    c-5.7,6.5-12,12.1-21.3,13.6c0-1-0.1-1.4,0-1.4c8.2-2.3,11.9-7.9,12.5-16.3c0.3-3.6,1.4-7.5,5.7-9.5c1.3-0.6,3.2-1.8,3-4.3    c-3-0.6-2.8,4.1-5.7,3.2c-0.7-5.7-0.2-11.2,5.6-14.6c6.3-3.7,6.4-3.6,9.7,3.1c1.2-1.1,2.9-1.8,2.5-3.9c0.5-0.5,1.2-0.7,1.6-1.2    c-0.5,0.4-1,0.8-1.4,1.4c-3.1-0.2-3.5-1.6-2.1-4.2c0.8-1.6,3-2.2,3-4.9c-5.5,1.6-10.4,8.6-16.8,2c5.7,0.6,9.4-2.3,13.1-6    c6-6,14.6-7.1,22.2-10.1c0.4-0.2,0.8-0.3,2.1-0.8c-8.1-0.8-13.7,4.3-20.5,5.3c0.6-3.3,4.4-2.6,5.8-5c-0.2-0.8-4.3,0-1.2-2.7    c4.4-3.7,9.6-4.9,15.1-5.9c-2.7,5.2-8.6,4.8-12.8,7.8c4.5,0.9,7.8-1.8,11.4-2.9c2.3-0.7,4.2-2.3,6.6-3c1-0.3,3.6-0.2,2.8-2.3    c-0.7-1.9-2.1,1.4-3.3,0.1c0.3-1.9,1.5-3.3,3.1-4.4c2,3,4.9,5.9-0.1,9.5c7-1.5,5.2-8.6,8.6-12.1c2.5,10.3-2.5,16.1-11.7,20.2    c11.6,0.1,13-4.8,16.3-18c2,1.5,3.5,3.5,5.2,4.7c3,2.2,5.5,6.7,11.4,4c2.1-1,7,2.9,9.3,6.4c0.9,1.4,1.5,2.9,2.9,3.8    c0.3-0.5,0.7-0.9,0.9-1.4c0.9-2.4-1.5-2.8-2.5-4c-1.4-1.7-3.2-3.1-4.4-5.2c7.6,4.9,15.4,9.5,21.2,16.8c-1.6-5.9-6-9-10.6-12    c-1.3-0.8-3.8-1.4-1.2-3.6c-1-0.5-0.3-2.3-2.1-2.4c-2.9-0.1-5.5-2.5-8.7-1.5c-1.4,0.4-2.7-0.7-3.7-1.7c-1-1.1-2-2.4-3.8-1.5    c-0.4,1.5,2.6,3,0.3,4.4c-2.1,1.3-2.7-1.2-3.8-2.2c-2.9-2.6-5.6-5.2-9.4-6.7c-5.7-2.2-6.2-5.1-3-11.7c1.7,4.6,6.3,5.6,9.2,8.7    c-3.2-1.4-6.3-2.8-9.5-4.3c7.1,6.9,17.1,7.2,25.2,11.6c-7.2-6-16.4-9.1-23.3-16.7c10.2,0.2,14.6,8.3,21.6,12.6    c-0.8-3.7-0.8-3.7,3.7-3.1c-2-2.8-2.2-3.3-3.7-2.7c-3.4,1.4-4.7-0.1-6.3-2.9c-2.2-3.7-6.2-6.2-9.9-7.6c-3-1.1-7.1,0.9-10.2,2.7    c-1.9,1.1-5.2,1.1-5.2,4.5c-0.4-0.1-0.8-0.3-1.2-0.5c0.4,0.3,0.9,0.4,1.4,0.4c1.4,3.8,2-1.8,3.7,1.1c0.4,2.9-2.6,6.6-2.7,11.3    c-3-11.6-11.5-16-22.2-17.5c3.8,1.8,7.7,3.5,11.4,5.4c4.1,2,5.9,6.8,3.8,9.7c-0.7-6.4-11.1-13.3-20.9-13.3    c5.7,3.6,12.4,3.8,16.8,9.7c-9.4-1-15.5-9.6-25.2-8.9c2.3,3.1,6.4,2,8.6,4.5c-5.3,2.2-9.2-2.7-14.6-2.8c1.1,1.7,3.3,0.7,3.6,2.9    c-3.1-0.8-6,1.7-9.2,0.9c9.6,1.8,19.6-0.1,29,3c-12.7-0.3-25.8,2.6-37.9-3.8c-2.1-1.1-5.4-2.4-4.9-4.6c0.5-2.1,3.9-2.1,6.4-2.3    c11.9-0.7,23.7-1.4,35.3-4.5c3-0.8,6.3-0.3,9.4-0.4c1.7-0.1,3.8,0.5,4.9-0.3c4.4-3.3,8-1.3,11.8,1.1c0.9,0.6,2.6,1.4,3,1    c5.5-4.4,8.9-0.1,12.7,2.8c1.1,0.9,2.5,2.2,3.7,2.2c4.8-0.1,9.1,1.9,13.6,2.8c6,1.2,12.4,2.2,18.7,1.3c1.2,1.5,3.2,1.1,4.8,1.9    c0.5,0.4,1.1,0.6,1.8,0.7c0.3,0.1,0.6,0.2,0.9,0.3c-0.3-0.3-0.7-0.2-1-0.3c-0.5-0.1-0.9-0.2-1.4-0.5c-1.5-1.1-3.7-0.7-5.2-2    c0.5-2,3.4-0.3,3.8-2.7c-2.9-1.5-6-3-9.4-4.6c-0.5,3.7,2.9,3.1,4.1,4.6c-4.8,1.2-6,0.9-10.5-1.7c-3.1-1.8-4.1-6.3-9.1-6.8    c2.3,3,4,5.3,6.5,8.6c-7.8-2.2-11.8-7.8-16.8-12.3c0.6,3.6,3.6,5.7,5.3,8.8c-5.5,0.9-9.4-2.3-12-6c-2-2.7-4.3-4.3-7.1-5.5    c-2.3-1-5.1-1.1-6.1-4.1c0-0.1-1.5-1.4-2.3-0.5c-0.3,0.4-0.1,1.8,0.3,2.2c1.6,1.4,2.6,3.5,4.8,4.2c0.6,0.2,2,0.7,1,2.1    c-0.7,1-1.7,1-2.7,0.7c-1.6-0.4-7.6-8.7-7.3-10.2c0.2-0.8,1.5-2.3,1.7-2.2c2.1,1,4.1-2.1,6.2,0c2,2.1,3.9,4.5,7.6,4.3    c-3.7-3.2-6.3-7.7-11.8-6.9c-3.1,0.5-5.2-0.1-8-2.2c-4.3-3.2-9.1,0.1-13.6,1.5c4.5,1.1,9.5-0.6,13.3,3.4c0.7,0.7,1.6,1.2,1.3,2.1    c-0.7,1.9-2.1,0.1-2.6-0.1c-7.8-3.1-15.5-2.1-23.4,0.1c-7.7,2.1-15.4,0-23.1-1c-1-0.1-2-1.1-3.2-0.9c-6.5,1.3-11.4-2.9-17-4.8    c-2.9-1-4.9-3.4-5.4-5.9c-0.5-2.5,2.7-3.4,4.7-4.2c2.6-1.1,2-3.1,2-5c0-1.1-0.8-1.9-1.8-2c-1.3-0.2-1,1-1.1,1.7    c-0.5,2.5-3.1,3.1-4.7,3.9c-2.3,1.1-4.5-1.1-4.1-3.4c0.5-3.4,1.6-7.4,5.6-8.1c16.5-3.2,31.6-10.7,47.6-15.1    c17.8-4.9,32.7-15.5,49.1-23.2c12-5.7,24.7-7.8,37.6-10.1c6.5-1.1,13.1,0.1,19.6-1.4c1.2-0.3,3.2,0.5,4.9-0.8    c1.1,3.4-2.3,2.9-3.2,4.7c0.3,1.2,3.3,0.1,2.7,2.5c-0.7,2.8-3.2,1.7-4.8,2.4c2.7,1.9,8.4,0.8,10.6-2.9c3.5-5.7,8.4-9,14.4-11.2    c3.4-1.2,6.8-1.5,10.5-1.5c4.4,0,8.8-1.6,12.9,1.7c0.5,0.4,2.6-0.9,3.8-1.7c0.9-0.5,1.6-1.2,2.6-0.1c0.6,0.6,1.2,1.2,0.6,2.1    c-1,1.4-1.6,1-2.1-0.4c-2,0.9-2.7,2.8-2.6,4.6c0.6,9.8-13.5,12.8-19.4,10c-0.7-0.3-1.2-1-2.1-1.1c-1.6-0.1-3.8,1.9-4.5-1.8    c-0.2-1.1-3.5,0.8-5.4,1.1c-1,0.2-2,0.9-3.3,0.8c0.9,1.9,3.6-0.1,3.6,1.9c0,2.5-3,2-4.2,3.5c1,0.9,2.5,1.9,1.7,3.7    c-0.8,1.9-3.4,2.7-8.7,3.1c1.7,1.7,1.7,1.7,3.2,1.7c5.3,0,10.4,0.3,14.8,3.9c0.6,0.5,1.9,0.3,2.7-0.6c1.1-1.3-0.2-1.9-0.8-2.1    c-1-0.4-1.8-0.8-2.1-1.9c2.3-1,4.3,1.7,6.6,0.6l-0.1-0.1c2.7,2.2,4.8,2.1,7.1-0.9c1.3-1.8,2.5-3.8,5.2-4.6    c1.6-0.5,3.9-2.3,2.7-5.1c-1.9,0.6-2.6,3.9-6.4,2.1c1.3-0.8,2.2-2,3.2-2c2.3-0.1,3.6-4.7,5.1-2.4c1.1,1.8,0,6.1-2.5,8    c-3.3,2.5-6.4,5.2-8.9,9c2.5-0.1,5.1,1.2,6.5-1.3c0.4-0.8-1.7-0.2-2.5-1.2c2,0.1,3.6-3.4,4.3-2c2.2,4.2,8.6,4.9,8.6,10.6    c0,3.2-0.5,6.4,0.2,9.4c0.7,2.9,0.8,5.8,0.8,9.3c-5-1.7-7.7-7.7-13.5-7.2c-1.8,0.3-3.7,0.7-5.6,0c1.9,1.2,3.8,0.2,5.6-0.1    c-0.1,0.3-0.3,0.8-0.2,0.9c7.7,4.3,14.2,10.1,20.3,16.5c-2.3-0.1-4.7,0.9-6.6-1.4c-5.2-6.5-12.8-9-20-12c-3.7-1.6-5.5-5.6-9.9-6.2    c2.2-3,5-2,7.4-2.3c-3.1-4.8-12.8-6.6-17.2-3.1c3.9,0.6,7.8-0.4,11.4,1.1c0.4,0.2,1,0,1,1.1c-0.9,0-1.8-0.1-2.7,0    c-1.8,0.2-4-2.7-5,0.1c-0.4,1,2,3.7,3.7,4.5c6.3,3.1,11.3,8.6,18.6,10.2c-4,2.2-8.9-0.2-12.9,2.8c6.4,0.4,11.4,6,18.6,3.9    c-3.4-1-6.2-2.7-9.5-3.1c3.3,0.1,7.8-3.4,9.1-1c3.3,6.1,12.6,8.9,10.1,18.4c-1.1-2.1-2.1-4-3.1-6c-3.4-6.8-6.9-8.5-14.2-6.6    c-3.6,0.9-7.2-1.7-11.4,0.4c6.1,0.2,11.4,2.1,17.2,2c4-0.1,4.4,4.7,6.4,8.1c-3.9-0.9-5.5-3.7-7.7-5.7c-1.1-1-2.4-1.8-3.4-0.6    c-0.9,1.1,0.9,1.8,1.4,2.7c2.8,5.3,9.3,5.3,13.1,9.3c0.6,0.6,1.1,1.3,1.1,1.9c0.2,4.1-0.4,8.1,0.8,12.2c-2.9-2.6-7.4-4.2-3.5-9.6    c1.4-1.9-0.3-4.1-3-5.4c-0.3,3.1,3.4,8.6-4.1,6.6c1.9,5.4,7,7.9,10.5,11.7c2.3,2.6,3.6,6.4,7.3,7.7c0.7,0.2,1.4,1.2,0.6,2    c-0.6,0.6-1.6,0.6-2.3,0.1c-3.3-2.7-7.2-4.6-9.5-8.7c-2.2-3.9-5.5-7.2-11.2-6.5c7.1,4.1,13.5,8.5,9.8,18c-1.3-0.5-2.5-1.4-2.5-3.4    c0.1-6.5-2.8-9.3-10.2-8.9c2.8,2.5,7.5,1.7,8.2,6.8c-1.8-1.5-3.2-2.7-4.9-4.2c2.9,9.4,10.6,14.2,17.5,21    c-10.1,0.2-13.9-9.1-21-12.2c-0.3,0.4-0.7,0.8-1,1.2c3.4,3.4,6.7,7,10.3,10.2c3.9,3.6,5.5,7.9,5.1,14.5c-2.7-3.9-7-5.1-7-10.1    c0-4.7-4.6-3.5-6.8-5.2c-3.5-2.7-9.9-0.2-11.1-6.4c-0.2-1.1-0.9-1.2-1.8-0.6c-5.7,3.9-10.9-0.2-16.3-1c-0.8-0.1-1.5-0.4-2.4-0.1    c1.9,2.9,4.6,3.8,7.8,3.7c6-0.1,11.6,0.3,16.4,4.8c1.3,1.2,3.5,2,5.7,2.3c3.1,0.5,6.5,1.3,6.9,5.6c0.2,2.1,2.5,2.7,3.7,4.1    c4,4.7,2.7,10.2,2.7,15.7c-1.3-5.9-3.2-11.2-10.4-12.2c6.6,3.9,7.3,10.9,9.6,17.3c-1,0.3-1.3-0.5-1.8-1c-0.4-0.5-0.8-1-1.1-1.6    c-3.8-8.6-10.2-14.3-19.3-16.9c-1.4-0.4-3-0.9-2.5-2.3c0.6-1.9,1.9,0.5,2.7,0.2c1.9-1,4.5-1,5.3-3.5c0.9-2.6-1.6-3.5-2.9-4.2    c-5.4-2.6-10.3-6.5-17.2-5.3c-6.8,1.2-13.2-1.6-19.7-3.6c4.2,4.2,9.8,4.4,15.1,5.7c0.6,0.2,1.3,0.1,2,0.1    c-1.4,1.8,0.1,1.5,1.2,2.2c3.2,2.1,1.8-1.6,2.8-2c1.9-0.8,3.7,0.2,5.6,0.6c4,0.8,6.9,3.1,9.9,5.6c-4.8,2.4-7.4-3.4-11.7-2.4    c0.4,1.9,2.1,2.5,3.4,3.6c-2.3,1.4-4.2-2.2-6.4,0c2.9,0.9,4.9,3.3,7.1,5.5c3.5,3.6,8.8,3.4,12.8,6.6c3.1,2.6,5.7,5.3,7.4,8.6    c2.8,5.4,5.9,10.7,7.9,16.5c1.2,3.7,5,5.6,7.5,8.8c-6.4-0.8-10.1-4-12.1-10.1c-1.5-4.7-3.1-9.6-7.8-12.4c-0.4-0.2-1.2-1.5-1.9-0.4    c-0.3,0.4-0.3,1.5,0.1,1.9c5.7,5.6,4.9,13.8,8.7,20.2c3.6,6.1,2,11.5-2,17c-3.1,4.3-3.2,9.8-0.5,13.8c3.2,4.7,7.5,9.1,14.2,9.5    c2.8,0.2,3.4-3,5.9-3.1c-0.8,3.7-6,6.9-9.8,5.8c-7.9-2.2-14.9-5.5-16.7-14.7c-0.3-1.3-0.6-2.2-2.6-1.4c-2.4,1-0.7,1.7-0.5,2.9    c2.1,8.5,9.3,12,15.8,16.2c2.2,1.5,4,3.1,6,6.1c-6.4-0.1-10.9-2.7-15.1-5.8c-4.8-3.6-7-8.9-9.7-13.9c-0.5-0.8,0.2-4-2.6-2.3    c-1.8,1.1,0,2.2,0.3,3.2c1.2,3.6,4.3,6.3,5.4,10.7c-5.8-3.1-9.6-7.1-11.1-13.3c-0.6-2.6-0.8-5.9-4.7-6.6    c1.6,9.5,4.5,17.9,14.4,22.1c-3,1.4-3,1.4-12.4-2.6c3.1,6.1,7.8,7.9,13.7,5.1c0.6-0.3,1.2-0.9,1.7-0.3c0.4,0.4,0.6,1.5,0.3,1.8    c-1.7,1.9-0.7,3.1,1.7,3.7c-1.8,1.4-3-0.8-4.1,0.7c2,1.2,3.5,3.2,6.2,3.4c1.4-3.2-3.3-5.3-1.4-8.3c4.6,2.1,5.4,6.6,6.7,11.4    c-13.3-0.6-20.6-10.7-30.2-18.6c2.3,8.2,8.5,11.6,14,15.8c-5,1.3-5.2,1.5-14.3-5.4c-4.8-3.7-9.1-8.2-12-13.7    c-0.4-0.7-0.2-2.8-2.1-1.8c-1.3,0.7-1.6,1.9-1.1,3.4c1.1,3.3,4.1,5,6,7.6c3.3,4.5,6.7,9.3,12.8,11c-5.6,3-6.6,2.7-11-2.5    c-2.2-2.5-3.5-5.9-6.6-7.6c-1.8,2.5,1.9,3.2,1.4,5.6c-3.6-2.1-5.4-6.4-9.6-7.5c-0.9,2.4,0.6,3.6,1.9,4.3c3.6,1.9,5.5,5.6,9,7.5    c0.7,0.4,2,0.7,1.1,2c-0.7,1-2,1.3-2.6,0.8c-4.1-3.7-10.1-5.3-12.7-10.8c-1.4-2.9-3.6-6.1,1.4-8.1c0.6-0.2,0.8-1,0.3-1.7    c-8.7,2.8-10.3,1.4-14.5-13.3c-2.7,5.1-0.3,12.7,5.7,15.8c2.8,1.4,3.7,3.4,4,5.8c0.5,4.8,3.8,7.7,7.4,12.1    c-9.9-3.4-14.5-9.8-16-19.4c-2.1,5.2-1.9,6.3,2.5,15.2c-6.3-1.1-6.4-1.3-11.8-12.5c-2,4.7-1.2,6.5,3.8,11.5    c3.3,3.3,7.9,2.9,12.8,4.7c-7,1.3-8.1,1.2-13.7-2.3c-3.9-2.4-6.5-5.6-8.3-10.1c-1.9-4.8,0.9-9,0.1-13.5c-6.7,6.5-5.4,14.9,3.8,25    c-4.3,0.3-5-0.5-13-14c-1.2,6.3,0.9,10.8,7.1,15.3c-6.2,1.7-9.2-4.5-14.5-6.3c0.8,3.8,2.9,5.5,5.3,7c-3,2.2-8.3,2.2-12.6,0    c-3.8-1.9-7.6-3.8-11.7-5.8c-0.5,2.6,1.4,3,2.7,4c-4.7,1.1-7.8-0.1-10.2-4.2c-2-3.5-6.1-5.3-7.9-9.3c-2-4.5-2.2-8.7,0.1-13.2    c-7.2,6.7-5.5,12.1,7.6,26.3c-8.8,2-19.9-7.5-19.4-17c2.4,0.2,2.1,5.6,6.2,2.8c-7.4-9-2.3-16.8,2.2-25c-4.1,0.6-4.4,5-6.2,7.5    c-1.6,2.3-1.9,5.8-2.3,8.8c-0.3,2.3-0.5,4.3-1.9,6.3c-2.5,3.7-0.2,7.3,2.2,9.9c2.7,3,6.1,5.3,9.2,7.9c-4.4,4.9-9.6,4.7-16.1-0.7    c-0.3-0.2-0.5-0.5-0.7-0.7c-1.8-1.1-2.9-4.9-5.4-2.7c-2,1.8,1.9,3.2,2.4,5.6c-4.4-1.2-8.2-2.8-11-6c-1.9-2.2-3.3-2.7-5.4-0.1    c-4.4,5.4-10.2,8.8-16.8,10.7c0.9-3,0.7-6.1-0.7-8.8c-2.3-4.5-2.1-8.6,2.2-11.3c3.8-2.4,5.8-6.1,9.1-9.7c-4.9,0.1-6.2,4.2-8.7,6.3    c-4.8,3.9-9.8,6.2-16.4,5.6c0.5-0.7,0.9-1.6,1.2-1.6c8.3,0.7,11.1-5.5,14.1-11.3c4-7.8,9.4-14.4,16.3-19.9    c0.9,1.9-0.9,2.4-1.7,3.2c-1,1.1-1.5,2.2-0.6,3.5c2.4,3.6,1.8,6.6-2.1,9.2c-1.3,0.9-2.7,1.6-4.3,2.5c2.4,0.6,4.4,0.8,6-1.2    c0.6-0.8,1.3-2.1,2.2,0.1c0.3,0.9,2.6,0.8,1,2.7c-1.2,1.5-1.8,3.5-3.1,4.9c-1.7,1.9-1.9,4.2-0.6,5.8c1.5,1.8,3.8,2.8,6.6,1.4    c-5.1-4-4.4-9.5,1.9-13.9c2.2-1.5,4-3.1,5.7-5.2c-2.9-1.2-3.9,1.6-5.8,2.2c-1.4,0.5-2.4,0.6-3.6-0.5c-1.3-1.3,0-2.2,0.5-3    c1.4-2.3,2.9-4.6,3-7.9c-1.7,0.5-2.5,2.7-4.3,1.7c0.1-6.2,5.3-8.6,9.2-11.9c0.5-0.4,1-0.8,1.5-1.1    C92.9,240.2,92.4,240.6,92.1,241.1z"
	pathData2 := "M107.1,148.7c-5.3,0.9-7.7-1.5-9.9-4c-2.1-2.4-4.7-2.4-5.5,0.2c-2.6,8.5-10,6.5-15.4,6.8    c-7.3,0.3-13,3.3-18.7,7.1c-1.2,0.8-2.4,2.7-4,1.2c-1.4-1.2-1-2.9,0.1-4.4c4.1-5.8,8.7-10.9,16-12.6c1-0.2,1.8-0.5,2.7-0.9    c4-2,7.1-0.9,9,3.1c0.9,1.8,1.8,3.2,4,2.7c1.9-0.4,3.1-1.6,3.4-3.7c0.4-2.7,1.6-3,4.5-2.5c5.1,0.8,10.4,0.4,15.5-1.3    c5.8-1.9,11.1,0.6,16.4,2.7c1.1,0.5,1.8,1.5,1.4,2.9c-0.3,1.5-1.6,0.7-2.4,1C118.4,149.3,111.7,146.5,107.1,148.7z"
	pathData3 := "M195.9,51.9c-6.3-1.9-10.6,2.1-15.8,5.8c6.4-15.1,7.3-15.8,22.2-19.8c5-1.3,9.9-3,14.8-5    c-1.4,1.2-3,2.2-4,3.7c-0.6,0.9-4.5,2.5,0.1,3.8c0.5,0.1-0.7,4.3-2.6,5.2c-2.3,1.1-3.2,3.7-5.6,4.4c4.4-0.4,6.5-4.8,10.7-5.5    c0.3,3.4-0.5,6.6-3.1,9.4c-0.6,0.7-1.5,1.9-2.4,0.8c-1-1.3,0.5-1.5,1.4-1.8c-1.4-2.7,2.7-2.9,2.1-5.9c-6,5.1-12,10.1-20.4,7.5    c0.4-1,2.9-0.6,2.3-2.7c3.8-0.9,6.1-4.3,8.1-5.3C201.8,47.6,199.3,50.4,195.9,51.9z"
	pathData4 := "M94.4,144.2c0.9,3.3,3.3,4.3,4.5,6.8c-2.5-0.2-5.2,1-6.6-0.6C90.7,148.6,94,147.3,94.4,144.2z"
	pathData5 := "M235.2,73.2c3-0.1,5.8-0.3,8.6-0.4C242.4,76.1,240.2,76.3,235.2,73.2z"
	pathData6 := "M82,142.1c2-0.1,3.5-0.4,5.1,0.4c-0.6,1.2-1.1,2.4-1.9,4.1C84,144.9,83.1,143.7,82,142.1z"
	pathData7 := "M239.8,80.9c1.3-1.3,2.1-3.4,5.4-2.5c-2.5,0.5-3,3.4-5.6,2.4C239.7,80.7,239.8,80.9,239.8,80.9z"
	pathData8 := "M134.1,103c3.1-3.4,9.1-1.2,11.1-6.6c-5.4-1-10.4-0.2-15.5,1.5c5.1-3.3,4.7-11.2,11.8-12.7c-0.1,3.5-4.4,4.9-4.3,9    c2.6-2.7,5.8-3.9,7.1-7.6c0.9-2.7,3.7-3,6.7-2.1c-5.4,1-4.6,6.7-8.1,9.3c1.6,0.3,2.9,1,3.2-1.3c0.3-2.4,1.9-4.1,3.7-5.6    c0.6-0.5,1.9-1.2,2-0.4c0.3,5.8,3.4,1.4,5.3,1.7c2.9,0.4,5.8,1.5,7.5,4c3,4.2,7.3,4.4,11.5,3.7c2.2-0.4,5.6,2.8,6.7-1.7    c5.1,4.3,3.7-2,5.4-3.2c1.5-1.1,2.9,1,4.3,1.4c2.3,0.8,2.9,3.5,2.5,5.9c-0.2,1.1-1.6,2.4-2.6,1.7c-2.1-1.8-4,0.9-6.1-0.2    c2.1-1.4,4-2.8,4.5-6.7c-4.8,6.2-10.4,7.6-16.9,7c-3.4-0.3-6.3-0.2-8.9,3.2c-3.9,5-10.2,5.3-15,1.1c-2-1.7-4-1.2-6-1.3    C140.8,102.9,137.6,103,134.1,103z"
	pathData9 := "M239.5,232.2c5.8,5.2,2.4,12.1,3.7,18.4c2.9-4.6,4.7-10.4,1.9-14.3c-3.1-4.4-1.2-9.4-4.4-13.9c-3.2-4.6-1.9-5.5-4.6-7    c5.2-3.2,5.2-3.2-0.4-7.5c-0.3-0.3-0.7-0.7-0.7-1.1c0.1-0.8,0.8-1.1,1.4-0.7c2.4,1.9,3.3-0.4,4.2-1.6c1.5-2.2-1.3-2.4-2.1-3.5    c-1.4-1.9-4.1-2.1-5.2-4.4c6.2-1.2,10.9,3.7,9.7,9.6c-1.1,5.2-2.1,11.7,0.4,15.6c2.7,4.3,3.7,8.4,5,13c1.7,6.3-0.9,11.8-2.6,17.4    c-0.3,1-1.1,2-2.6,1.5c-1.3-0.4-2.4,2-4,0.7c-5.1,8.3-5.7,0.6-7.2-2.2c-1.8-3.5-2.9-7.4-5.1-10.8c5,3,6.6,8.4,9.3,13    c1.8-2.3-0.2-4.4-0.4-6.6c-0.2-2.2-0.1-5.1-1.4-6.5c-1.9-2.1-1-4.3-1.8-6.4c2.3,4.8,5.8,9.1,6.2,14.6c0.1,1,0.5,2.1,1.7,1.8    c1.3-0.3,0.4-1.5,0.5-2.2c1.3-5.1,0-9.8-2.1-14.3C238.8,234.3,238.6,233.1,239.5,232.2z"
	pathData10 := "M190.1,175.1c-2.5,4.7-5.7,1-8.1,0.1c-1.8-0.7-3.8-1.3-6.7-2.9c2.1,4.7,5.1,7.4,7.8,9.5c4.6,3.4,7.4,8.1,11.2,12.1    c0.5,0.5,0.9,1.2,0.4,1.7c-0.8,0.7-2.2,0.8-2.7-0.1c-3.8-7.6-13.1-10.2-16.5-18.3c-1.1-2.6-3.9-3.5-6.4-4.1    c-2.5-0.6-4.5-1.6-6.1-3.7c4-2.2,5.9,4.1,9.8,2.2c0-3.2-4.4-2.4-4.5-5.8c7.6,2.4,13.5,8.8,21.8,9.4L190.1,175.1z"
	pathData11 := "M241.2,63.3c5.2,3.4,8.6-3.2,14.1-2.1c-3.2,2.5-6.1,4.1-9.5,4.7c-3.2,0.6-6.5,2.2-9.3-1.6c-0.7-0.9-3.4-0.4-5.3-0.8    c3.6-0.5,5-3.9,7-5.9c3.6-3.8,7.5-2.2,11.3-2.5c1.8-0.2,3,1.4,3.4,3c0.7,2.6-1.9,1.6-2.8,2.3C247.5,62.4,243.7,61.1,241.2,63.3z"
	pathData12 := "M190.3,175.3c-0.7-0.9-1.4-1.8-2.2-2.6c-1-1-1.9-2.5-0.6-3.3c1.7-1.1,1.2,1.3,2.1,1.8c5.6,3.1,11.6,5.1,17.7,6.8    c2.1,0.6,4.1,0.9,5.6,2.6c-1.8,1-3.7,1.9-6,3.2c4.2,0.5,7,2.8,8.7,5.8c1.5,2.6,4.8,4.7,4.2,8.9c-4.7-2.1-3.9-10.7-11.8-8.1    c2.7-1.7,3.3-3.1,0.2-4c-1.6-0.5-4.5,0.7-4.7-1.4c-0.2-1.8,2.2-3.4,5-4.3c-7-1-11.8-6.2-18.4-5.4    C190.1,175.1,190.3,175.3,190.3,175.3z"
	pathData13 := "M220.5,186.8c-1,2.3,0.8,2.8,1.5,3.7c3.2,4.1,4.2,11.8,1.9,16c-0.5,0.9-1.2,1.8-2.3,1.4c-0.3-0.1-0.7-1.4-0.5-1.7    c5.2-7.2-1-12.4-3.7-17.9c-0.7-1.5,1-3.2-1-4c1.3-0.2,2.5-0.5,3.8,0.4c3,2.2,7.8,2.1,9.1,7.8C225.9,190.2,223.2,188.5,220.5,186.8    z"
	pathData14 := "M184,181.2c5.7-5.6,10.6,0.1,15.7,0.2c0.4,0,1.4,1.6,1.2,1.8c-3.8,4.8,1.5,8.5,1.6,12.7c0,1.2,1.4,2.6-0.8,3    c-2.3,0.4-1.8-1.4-1.7-2.6c0.2-2-1.4-2.8-2.3-4.1c-1.2-1.8-3-4.1-1.2-6.2c3.6-4-1-3.7-2.2-3.8C191.1,181.7,187.9,180,184,181.2z"
	pathData15 := "M196.5,266c-6.6-0.2-6.3-5.7-7.3-9.4c-1.1-4.1,2.3-7.9,2-12.7c1.6,2.2,0,4.4,1.7,6.3c1.4-1.2,2.3-2.6,2.2-4.5    c-0.1-1.8,1-2,2.2-1.6c1.2,0.4,2.5,1.2,1.3,2.9c-0.7,1.1-1.4,2.1-2.2,3.2c-1.8,2.3-6.3,2.1-4.7,7.7    C192.6,261.4,195.2,262.9,196.5,266z"
	pathData16 := "M253.5,96.5c-1.7-1.1-3.2-3-5.8-2.7c-0.5,0.1-2.2-0.3-2.3-1.7c2.1,0,0-2.9,2.2-3.1c1.6-0.1,1.4,1.2,1.9,1.7    c3.1,2.7,7.2,3.9,10.2,7.3c-3.2-0.5-6,2.9-8.5,1.4c-2.6-1.5-5.1-2.5-8.5-2.6C246.7,92.5,250.2,99.5,253.5,96.5z"
	pathData17 := "M210,214.2c6.9-0.4,9.4,2.7,6.7,8.4c-1,2.2-0.1,4.5-1.5,6.6c-0.7,1-0.7,2.1-2.2,1.8c-1.5-0.3-0.9-1.5-1-2.4    c-0.1-1.1-0.5-2.7,0.1-3.4C216,220.7,213,217.5,210,214.2z"
	pathData18 := "M221.4,160.5c4.9,2.1,10.6-1.2,15.5,2.3c1.6,1.1,4.9,0.8,6.6-1.8c0.4-0.6,1.2-1.5,2-0.8c0.9,0.7,0.8,2,0.1,2.7    c-1.7,1.7-2.8,3.8-6.5,2.4C233.5,163.1,227,163.5,221.4,160.5z"
	pathData19 := "M160.5,113.1c-2.7,0-5.4,0-8.9,0c3.3,3.1,6.7,3,10.3,2.9c-1.6,3.3-4.2-0.7-5.6,1.9c0,0.8,1.4,0.9,2.3,1.6    c-2.1,0.9-3.2,0.3-4.8-1.2c-3.8-3.5-5.7-8.4-9.3-12.3C149.5,109,154.1,112.9,160.5,113.1z"
	pathData20 := "M250.3,190.9c-3.5,0.4-6.6-0.5-9.7-3.9c3.4,0.9,6.2,1.2,8.2-1.4c-0.8-2.5-4.4-1.7-5.6-4.4c3.5-0.4,6.4-0.3,8,3.5    c0.7,1.7,1.9,2.7-1,3.5C246.9,188.9,250.9,189.8,250.3,190.9z"
	pathData21 := "M180.9,211.5c5.9-0.5,11.5,3.6,12,8c0.1,0.8,0,1.9-0.4,2.3c-0.7,0.6-1.8-0.1-2-0.9c-0.4-2.5-1.9-4.1-3.7-5.5    C185.3,214.2,183.6,213.2,180.9,211.5z"
	pathData22 := "M214.3,254.9c3,5.9,3.6,12.3,8.9,16.3C218,271.4,213.7,263.7,214.3,254.9z"
	pathData23 := "M218.1,215.9c0.8-6.1-3.4-9-7-12.4c4.6,0.7,7.5,3.6,9.4,7.6C221.5,213,221.3,213.1,218.1,215.9z"
	pathData24 := "M156.3,184.3c4.6,2,7.7,5.5,10.6,9.4c0.4,0.6,0.4,1.7-0.4,2.2c-0.3,0.2-1.2,0-1.3-0.3    C162.9,191.2,158.2,188.9,156.3,184.3z"
	pathData25 := "M160.5,126.7c-7.1-0.4-13-2.5-15.4-10.1c-0.2-0.5-1.3-0.8-0.4-1.4c0.6-0.4,2-0.6,2.1,0.3    C148.3,123.1,154.9,124.2,160.5,126.7z"
	pathData26 := "M110.4,255c-2.4-5.7,3-14,9-14.7C115.7,244.7,110.4,248.1,110.4,255z"
	pathData27 := "M199.7,234c3.5,0.7,4.3,3.6,5.2,6.4c0.6,1.9-0.8,3-2.1,4.1c-0.7,0.6-1.8,0.8-2.2,0.1c-0.7-1.1,1-0.5,1.2-1.3    C202.7,239.9,199.3,237.7,199.7,234z"
	pathData28 := "M180.5,202.1c3.9,2.5,9.4,2.1,11.6,7.1c0.5,1.2,1.8,2.7,0.4,3.4c-1.9,1.1-1.8-1.5-2.3-2.3c-1.5-2.5-3.6-4.5-6.2-5.3    C182.2,204.6,181.2,203.7,180.5,202.1z"
	pathData29 := "M138.2,284.8c-5.1,0.2-8.7-1.5-11.7-4.8c-0.7-0.8-0.9-2-0.1-2.7c1-0.9,1.9,0.2,2.3,1C130.5,282.1,135.3,281.6,138.2,284.8    z"
	pathData30 := "M251.5,86.2c4.8-1.9,6,2.9,8.7,4.8c1.8,1.2,1.2,3.5-0.6,4.8C256.8,92.7,256.9,87.1,251.5,86.2z"
	pathData31 := "M197.3,168.8c-6.4,0.5-11.5-2.2-17-3.1c-0.7-0.1-1.5-0.9-1.3-1.9c0.2-1.1,1.6-1.1,2-0.7    C185.5,167.2,191.7,166.1,197.3,168.8z"
	pathData32 := "M248.8,227.4c4.1,4.8,3.3,10.1,2.8,15.3C250.7,237.6,249.8,232.5,248.8,227.4z"
	pathData33 := "M188.4,232.1c-3.1-1.4-2.5-5.3-5.2-6.6c-0.4-0.2-0.2-1.5,0.6-1.4c0.8,0.1,1-1.6,2.3-0.2    C188.2,226.3,190.2,228.4,188.4,232.1z"
	pathData34 := "M234.9,154.5c5.4-2.2,10.8-1.9,13.8-7.3C248.4,152.6,243.4,155.2,234.9,154.5z"
	pathData35 := "M71.8,282.9c0.5,5.1-0.4,9-3.9,12.1C68.5,291,69,287,71.8,282.9z"
	pathData36 := "M184.5,236.1c1.8-0.2,2.8,1.3,3.2,2.4c0.5,1.5,0.2,3.2,0.1,4.9c0,0.2-1.2,0.7-1.7,0.5c-2.2-0.6-0.7-2.5-1.1-3.7    C184.6,238.9,183.2,237.7,184.5,236.1z"
	pathData37 := "M261.8,107.8c-2.8-1.2-3.5-3.9-5.3-5.8c-1.2-1.3-0.1-2.1,0.8-2.7c1.4-0.8,1.5,0.3,1.7,1.3    C259.5,103.1,262.6,104.4,261.8,107.8z"
	pathData38 := "M241.5,144.7c4.2-1.6,7.6-4.9,13.4-6.1C250.5,143.2,245.6,143.2,241.5,144.7z"
	pathData39 := "M180.3,132.9c-6.4,0.6-10.6-2.5-14.8-5.8C170.2,128.9,174.9,130.8,180.3,132.9z"
	pathData40 := "M164.7,205.1c0-3.6-2.5-5.5-5.1-7.8C165.2,197.8,166.6,199.9,164.7,205.1z"
	pathData41 := "M218.5,165.4c-4.5,1.3-8-0.1-10.3-3.9C211.8,162.3,214.6,165.4,218.5,165.4z"
	pathData42 := "M156.6,234c1.4,2.7-1.2,4.2-1.7,6.4c-0.2,0.9-1.4,0.8-2.1,0.6c-1.1-0.4-0.8-1.2-0.2-2C153.8,237.3,153.6,234.4,156.6,234z    "
	pathData43 := "M99.4,214.2c-2,3.1-4.4,5.6-8.7,6.6C92.4,216.4,96.8,216.5,99.4,214.2z"
	pathData44 := "M173.9,123.9c4.3,0.5,7.5,2.7,11.3,2.2c0.2,0,0.7,0.7,0.6,1c-0.1,0.3-0.8,0.8-1,0.7C181.7,125.7,176.8,128.8,173.9,123.9z    "
	pathData45 := "M99.5,132.1c4.4,3,8.2-1,12.2-0.8C107.8,134.9,103.7,134.7,99.5,132.1z"
	pathData46 := "M249.5,102.5c-2.7,3.5-5.5,1.9-8.3-0.1C244,99.2,246.7,102.7,249.5,102.5C249.4,102.5,249.5,102.5,249.5,102.5z"
	pathData47 := "M247,134.6c1.7-2.4,2.8-4.6,0-8.3C253.4,131.4,253.3,131.6,247,134.6z"
	pathData48 := "M163.9,217.1c-0.4-2.1-2.1-3.9-1.1-6c0.3-0.5,1.9-0.2,2.6,0.8C167.1,214.3,164.9,215.4,163.9,217.1z"
	pathData49 := "M197.4,156c3.2,0.5,6.3,0.9,9.9,1.5C204.6,159.2,204.6,159.2,197.4,156z"
	pathData50 := "M232.8,127.2c3.6,0,6.7,0,11.1,0C239.5,128.4,236.4,130.7,232.8,127.2z"
	pathData51 := "M212.9,151.3c-3.3,2.9-6.2-0.5-9.4,0C206.6,151.3,209.8,151.3,212.9,151.3z"
	pathData52 := "M87.3,225.8c1-4.4,4.5-1.2,7-3.6C92.5,226.1,90.2,226.3,87.3,225.8z"
	pathData53 := "M220.5,133c3.7,0.4,6.7-1,9.6,1.6C226.7,135.5,224,134.8,220.5,133z"
	pathData54 := "M187.7,123.6c2.7-1.1,5.2,0,7.2-1.5c0.4-0.3,1.8-0.4,2,0.7c0.3,1.5-1.4,1.6-2,1.6C192.7,124.4,190.6,124,187.7,123.6z"
	pathData55 := "M112.9,199.5c2.7-0.9,5.5-1.8,8.2-2.7C119,199.5,116,199.6,112.9,199.5z"
	pathData56 := "M194,163c2.3,0,4.6,0,6.9,0c-0.1,0.5-0.1,1-0.2,1.6c-2.4,0.8-4.7,1-6.6-1.7L194,163z"
	pathData57 := "M223.7,97.6c-3.2,1.2-4.5-0.6-5.9-2.8c1.9-1.7,2.4,0.2,3.3,0.8C222,96.2,222.8,96.9,223.7,97.6z"
	pathData58 := "M149,138.4c3.2,1.6,4,3.6,3.6,6.3C152.2,142.5,149.1,142,149,138.4z"
	pathData59 := "M125.6,192.6c0.1,0.3,0.1,0.6,0.2,0.9c-1.8,0.2-2.4,3.4-4.8,2c-0.4-0.2-0.5-0.7-0.8-1.1c1.7-1,3.5-1.4,5.5-1.6    L125.6,192.6z"
	pathData60 := "M90,281.7c-0.3,1.4-1,2.4-2.5,2.3c-0.4,0-1.3,0.4-1-0.8c0.4-1.6,1.5-2.4,2.9-2.7C90.3,280.2,89.7,281.2,90,281.7z"
	pathData61 := "M216,153.6c4.2-0.8,6.9-0.7,9.6-0.1C222.9,154.2,220.2,154.1,216,153.6z"
	pathData62 := "M226.3,146.5c2.2,0.1,3.9-2.1,6.6-1.2C231,148,228.5,146.7,226.3,146.5z"
	pathData63 := "M221.4,181.6c-2.9,0.6-5.4-0.1-7.1,2.1C215.9,180.7,215.9,180.7,221.4,181.6z"
	pathData64 := "M118.7,154.6c3.1-1.4,3.8,0.7,5.2,1.8C121.3,158.2,120.6,155.5,118.7,154.6z"
	pathData65 := "M130.3,159.7c-2.5,1.1-3.9-0.4-5.2-2.3c2.7-1.3,3.6,1.1,5,2.4C130.2,159.8,130.3,159.7,130.3,159.7z"
	pathData66 := "M83.4,233c-1.1-2.2,0.5-3.5,1.3-4.9c0-0.1,1.2,0.3,1.2,0.4C85.7,230.3,84.8,231.7,83.4,233z"
	pathData67 := "M92,247.2c-1.2,1.5-2.1,3.3-4.8,2.5C88.3,248.1,89.1,246.6,92,247.2z"
	pathData68 := "M188.9,161.1c2.1-0.3,3.8,0.2,5.1,1.9c0,0,0.1-0.1,0.1-0.2c-1.7-0.6-4.7,2.2-5.1-1.9C189,160.9,188.9,161.1,188.9,161.1z"
	pathData69 := "M99.1,208.1c0.3-3.1,2.3-3.1,4.6-2.7c-1.8,0.4-2.4,3.2-4.8,2.6C98.9,207.9,99.1,208.1,99.1,208.1z"
	pathData70 := "M130.5,105.9c3.3-2.1,4.9,0.6,7.9,0.1C135.2,108.1,133.5,105.5,130.5,105.9z"
	pathData71 := "M168.3,82.3c-2.4-0.9-4.1-1.2-5.6-3.7C165.1,79.6,166.8,79.8,168.3,82.3z"
	pathData72 := "M125.7,192.7c1.7-1.3,4-0.3,5.8-1.4c-1.4,2.6-3.7,1.7-5.9,1.3C125.6,192.6,125.7,192.7,125.7,192.7z"
	pathData73 := "M105.8,196.4c1.3-2.7,3.1-2.7,5.4-3.8C109.9,195.3,108,195.3,105.8,196.4z"
	pathData74 := "M178.8,146.6c-2.4,1.1-3.5-1.5-5.4-0.9C175.5,144.3,177.1,145.4,178.8,146.6z"
	pathData75 := "M92.1,241.1c-0.5-2.9,2.4-1.8,3.3-3c-0.5,1.8-1.3,3.2-3.4,2.9C91.9,240.9,92.1,241.1,92.1,241.1z"
	pathData76 := "M165.2,142.4c2.1-1.2,3.1,0.6,4.6,1.1C167.8,144.8,166.8,143.1,165.2,142.4z"
	pathData77 := "M195.9,51.9c3-3.9,7.8-5.5,11.4-9.6c-0.8,4.2-7.6,9.4-11.5,9.4C195.7,51.7,195.9,51.9,195.9,51.9z"
	pathData78 := "M188.2,48.5c0.4-2.1,2.2-2,3.5-2.8c1.3-0.9,2.7-1.6,4.4-2.6C195.7,48.9,190.6,46.5,188.2,48.5z"
	pathData79 := "M156.4,92c3.7,0,7.9,3.2,7.6,6.1c-0.5,4.2-4.1,5.3-7.4,5.8c-3.8,0.6-6.5-2.3-6.6-6    C149.9,93.7,151.7,92,156.4,92z"
	pathData80 := "M152.1,96.4c1.4-0.3,3.5,4.4,3.8-0.7c0.1-2.2,1.5-2.2,2.6-1.4c1.2,0.8,4,0.6,3.1,3.4c-0.9,2.5-2.7,4.1-5.3,4.2    C153.8,102.1,152.4,100.1,152.1,96.4z"
	canvas.Scale(0.8)
	canvas.Translate(30, 30)
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+hexcode)
	canvas.Path(pathData5, "fill:"+hexcode)
	canvas.Path(pathData6, "fill:"+hexcode)
	canvas.Path(pathData7, "fill:"+hexcode)
	canvas.Path(pathData8, "fill:"+lineColor)
	canvas.Path(pathData9, "fill:"+lineColor)
	canvas.Path(pathData10, "fill:"+lineColor)
	canvas.Path(pathData11, "fill:"+lineColor)
	canvas.Path(pathData12, "fill:"+lineColor)
	canvas.Path(pathData13, "fill:"+lineColor)
	canvas.Path(pathData14, "fill:"+lineColor)
	canvas.Path(pathData15, "fill:"+lineColor)
	canvas.Path(pathData16, "fill:"+lineColor)
	canvas.Path(pathData17, "fill:"+lineColor)
	canvas.Path(pathData18, "fill:"+lineColor)
	canvas.Path(pathData19, "fill:"+lineColor)
	canvas.Path(pathData20, "fill:"+lineColor)
	canvas.Path(pathData21, "fill:"+lineColor)
	canvas.Path(pathData22, "fill:"+lineColor)
	canvas.Path(pathData23, "fill:"+lineColor)
	canvas.Path(pathData24, "fill:"+lineColor)
	canvas.Path(pathData25, "fill:"+lineColor)
	canvas.Path(pathData26, "fill:"+lineColor)
	canvas.Path(pathData27, "fill:"+lineColor)
	canvas.Path(pathData28, "fill:"+lineColor)
	canvas.Path(pathData29, "fill:"+lineColor)
	canvas.Path(pathData30, "fill:"+lineColor)
	canvas.Path(pathData31, "fill:"+lineColor)
	canvas.Path(pathData32, "fill:"+lineColor)
	canvas.Path(pathData33, "fill:"+lineColor)
	canvas.Path(pathData34, "fill:"+lineColor)
	canvas.Path(pathData35, "fill:"+lineColor)
	canvas.Path(pathData36, "fill:"+lineColor)
	canvas.Path(pathData37, "fill:"+lineColor)
	canvas.Path(pathData38, "fill:"+lineColor)
	canvas.Path(pathData39, "fill:"+lineColor)
	canvas.Path(pathData40, "fill:"+lineColor)
	canvas.Path(pathData41, "fill:"+lineColor)
	canvas.Path(pathData42, "fill:"+lineColor)
	canvas.Path(pathData43, "fill:"+lineColor)
	canvas.Path(pathData44, "fill:"+lineColor)
	canvas.Path(pathData45, "fill:"+lineColor)
	canvas.Path(pathData46, "fill:"+lineColor)
	canvas.Path(pathData47, "fill:"+lineColor)
	canvas.Path(pathData48, "fill:"+lineColor)
	canvas.Path(pathData49, "fill:"+lineColor)
	canvas.Path(pathData50, "fill:"+lineColor)
	canvas.Path(pathData51, "fill:"+lineColor)
	canvas.Path(pathData52, "fill:"+lineColor)
	canvas.Path(pathData53, "fill:"+lineColor)
	canvas.Path(pathData54, "fill:"+lineColor)
	canvas.Path(pathData55, "fill:"+lineColor)
	canvas.Path(pathData56, "fill:"+lineColor)
	canvas.Path(pathData57, "fill:"+lineColor)
	canvas.Path(pathData58, "fill:"+lineColor)
	canvas.Path(pathData59, "fill:"+lineColor)
	canvas.Path(pathData60, "fill:"+lineColor)
	canvas.Path(pathData61, "fill:"+lineColor)
	canvas.Path(pathData62, "fill:"+lineColor)
	canvas.Path(pathData63, "fill:"+lineColor)
	canvas.Path(pathData64, "fill:"+lineColor)
	canvas.Path(pathData65, "fill:"+lineColor)
	canvas.Path(pathData66, "fill:"+lineColor)
	canvas.Path(pathData67, "fill:"+lineColor)
	canvas.Path(pathData68, "fill:"+lineColor)
	canvas.Path(pathData69, "fill:"+lineColor)
	canvas.Path(pathData70, "fill:"+lineColor)
	canvas.Path(pathData71, "fill:"+lineColor)
	canvas.Path(pathData72, "fill:"+lineColor)
	canvas.Path(pathData73, "fill:"+lineColor)
	canvas.Path(pathData74, "fill:"+lineColor)
	canvas.Path(pathData75, "fill:"+lineColor)
	canvas.Path(pathData76, "fill:"+lineColor)
	canvas.Path(pathData77, "fill:"+lineColor)
	canvas.Path(pathData78, "fill:"+lineColor)
	canvas.Path(pathData79, "fill:"+hexcode)
	canvas.Path(pathData80, "fill:"+lineColor)
	canvas.Gend()
	canvas.Gend()
}