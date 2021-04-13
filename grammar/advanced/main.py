from PIL import Image as ImagePIL, ImageFont, ImageDraw
import base64
import io
import sys

def main(argv):
    msg="success"
    try:
        base64Str=argv[0]
        savePath=argv[1]
        xdpi=300.0
        ydpi=300.0
        argvLen=len(argv)
        if argvLen>2:
            xdpi=float(argv[2])
        if argvLen>3:
            xdpi=float(argv[3])
        if base64Str=="":
            msg="base64Str is empty"
            return msg
        if savePath=="":
            savePath="savePath is empty"
            return msg
        f = open(base64Str)
        base64Str=f.read()
        f.close()
        if base64Str.find(";base64,")>-1:
            base64Str=base64Str[base64Str.find(";base64,")+8:len(base64Str)]
        imgdata = base64.b64decode(base64Str)
        image=io.BytesIO(imgdata)
        img =ImagePIL.open(image)
        img=img.convert('RGB')
        img.save(savePath,dpi=(xdpi,ydpi))
        img.close()
    except Exception as result:
        msg=result
    finally:
        return msg

if __name__ == '__main__':
    msg=main(sys.argv[1:])
    print(msg)
