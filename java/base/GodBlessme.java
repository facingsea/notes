package code;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/**
 * Created by wangzhf on 2016/3/12.
 */
public class GodBlessme {
    private static final int SHUANG_SE_QIU = 0;
    private static final int DA_LE_TOU = 1;
    private static final int RED_MAX_NUM = 33;
    private static final int BLUE_MAX_NUM = 16;
    private static final int RED_COUNT = 6;
    private static final int BLUE_COUNT = 1;

    public static void main(String[] args) {
        //runShuangSeQiu();
        runDaLeTou();
    }

    private static void runDaLeTou(){
        List<Integer> list = new ArrayList<Integer>();
        List<Integer> list2 = new ArrayList<Integer>();
        for(int i = 0; i < 10; i++){
            // red
            int j = 0;
            while (j < RED_COUNT - 1){
                int red = (int)(Math.random() * 35);
                if(list.contains(red) || red == 0){
                    //j --;
                    continue;
                }else{
                    list.add(red);
                    j++;
                }
            }
            list.sort(new Comparator<Integer>() {
                @Override
                public int compare(Integer o1, Integer o2) {
                    return o1 - o2;
                }
            });
            for(int val : list){
                if(val >= 10){
                    System.out.print(val + "  ");
                }else{
                    System.out.print("0" + val + "  ");
                }

            }
            System.out.print(" , ");
            // blue
            int k = 0;
            while (k < 2){
                int blue = (int)(Math.random() * 12);
                if(list2.contains(blue) || blue == 0){
                    //j --;
                    continue;
                }else{
                    list2.add(blue);
                    k++;
                }
            }
            list2.sort(new Comparator<Integer>() {
                @Override
                public int compare(Integer o1, Integer o2) {
                    return o1 - o2;
                }
            });
            for(int val : list2){
                if(val >= 10){
                    System.out.print(val + "  ");
                }else{
                    System.out.print("0" + val + "  ");
                }

            }

            System.out.println("");
            list.clear();
            list2.clear();
        }
    }

    private static void runShuangSeQiu(){
        List<Integer> list = new ArrayList<Integer>();
        for(int i = 0; i < 10; i++){
            // red
            int j = 0;
            while (j < RED_COUNT){
                int red = (int)(Math.random() * RED_MAX_NUM);
                if(list.contains(red) || red == 0){
                    //j --;
                    continue;
                }else{
                    list.add(red);
                    j++;
                }
            }
            list.sort(new Comparator<Integer>() {
                @Override
                public int compare(Integer o1, Integer o2) {
                    return o1 - o2;
                }
            });
            for(int val : list){
                if(val >= 10){
                    System.out.print(val + "  ");
                }else{
                    System.out.print("0" + val + "  ");
                }

            }
            // blue
            int blue = 0;
            while(true){
                blue = (int)(Math.random() * BLUE_MAX_NUM);
                if(blue == 0){
                    continue;
                }else{
                    break;
                }
            }
            System.out.print(", " + blue);
            System.out.println("");
            list.clear();
        }
    }
}
