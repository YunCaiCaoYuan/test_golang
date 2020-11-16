import java.util.regex.*;

class RegexExample1{
    public static void main(String args[]){
        String content = "I am noob " + "from runoob.com.";
        String pattern = ".*runoob.*";
        boolean isMatch = Pattern.matches(pattern, content);
        System.out.println("is include 'runoob' " + isMatch);

        content = "34567";
        pattern = "(?:0(?=1)|1(?=2)|2(?=3)|3(?=4)|4(?=5)|5(?=6)|6(?=7)|7(?=8)|8(?=9)){5,}\\d";
        isMatch = Pattern.matches(pattern, content);
        System.out.println("is include '123456789' " + isMatch);
    }
}