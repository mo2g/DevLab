
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

@Slf4j
public class IdCertificationUtil {

    /** 性别表示的值 */
    private static final int MALE_SEX_INT = 1;
    private static final String MALE_SEX_STRING = "男";
    private static final int FEMALE_SEX_INT = 2;
    private static final String FEMALE_SEX_STRING = "女";
    /** 第一代居民身份证长度 */
    private static final int FIRST_ID_CARD_LENGTH = 15;
    /** 第一代居民身份证 年份值 */
    private static final String FIRST_ID_CARD_YEAR = "19";

    /** 第二代居民身份证长度 */
    private static final int SECOND_ID_CARD_LENGTH = 18;
    /** 第二代居民身份证 校验码的模 */
    private static final int SECOND_ID_CARD_CHECK_MOD = 11;
    /** 性别 map 中的 key 值 */
    public static final String SEX_BY_INT_MAP_KEY = "sex_by_int";
    public static final String SEX_BY_STRING_MAP_KEY = "sex_by_string";
    private static SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd");


    /**
     * @desc 通过身份证获取性别
     * @auth llp
     * @date 2022/7/15 16:10
     * @param idCard 身份证
     * @return java.util.Map<java.lang.String,java.lang.Object>
     */
    public static Map<String, Object> getSexFromIdCard(String idCard){

        Map<String, Object> sexMap = new HashMap<>();
        // 默认值
        int sexInt = 0;
        String sexStr = "未知";
        // 15 位身份证
        if (idCard.length() == FIRST_ID_CARD_LENGTH){
            String sex = idCard.substring(14, 15);
            // 偶数表示女性，奇数表示男性
            if (Integer.parseInt(sex) % 2 == 0){
                sexInt = FEMALE_SEX_INT;
                sexStr = FEMALE_SEX_STRING;
            }else {
                sexInt = MALE_SEX_INT;
                sexStr = MALE_SEX_STRING;
            }
        }
        // 18 位身份证
        if (idCard.length() == SECOND_ID_CARD_LENGTH){
            String sex = idCard.substring(16, 17);
            log.info("sex",sex);
            // 偶数表示女性，奇数表示男性
            if (Integer.parseInt(sex) % 2 == 0){
                sexInt = FEMALE_SEX_INT;
                sexStr = FEMALE_SEX_STRING;
            }else {
                sexInt = MALE_SEX_INT;
                sexStr = MALE_SEX_STRING;
            }
        }
        // 结果
        sexMap.put(SEX_BY_INT_MAP_KEY, sexInt);
        sexMap.put(SEX_BY_STRING_MAP_KEY, sexStr);

        return sexMap;
    }

    /**
     * 根据身份证号获取年龄
     * @param IDCard
     * @return
     */
    public static Integer getAge(String IDCard){
        Integer age = 0;
        Date date = new Date();
        if (StringUtils.isNotBlank(IDCard)&& isValid(IDCard)){
            //15位身份证号
            if (IDCard.length() == FIRST_ID_CARD_LENGTH){
                // 身份证上的年份(15位身份证为1980年前的)
                String uyear = "19" + IDCard.substring(6, 8);
                // 身份证上的月份
                String uyue = IDCard.substring(8, 10);
                // 当前年份
                String fyear = format.format(date).substring(0, 4);
                // 当前月份
                String fyue = format.format(date).substring(5, 7);
                if (Integer.parseInt(uyue) <= Integer.parseInt(fyue)) {
                    age = Integer.parseInt(fyear) - Integer.parseInt(uyear) + 1;
                    // 当前用户还没过生
                } else {
                    age = Integer.parseInt(fyear) - Integer.parseInt(uyear);
                }
                //18位身份证号
            }else if(IDCard.length() == SECOND_ID_CARD_LENGTH){
                // 身份证上的年份
                String year = IDCard.substring(6).substring(0, 4);
                // 身份证上的月份
                String yue = IDCard.substring(10).substring(0, 2);
                // 当前年份
                String fyear = format.format(date).substring(0, 4);
                // 当前月份
                String fyue = format.format(date).substring(5, 7);
                // 当前月份大于用户出身的月份表示已过生日
                if (Integer.parseInt(yue) <= Integer.parseInt(fyue)) {
                    age = Integer.parseInt(fyear) - Integer.parseInt(year) + 1;
                    // 当前用户还没过生日
                } else {
                    age = Integer.parseInt(fyear) - Integer.parseInt(year);
                }
            }
        }
        return age;
    }

    /**
     * 身份证验证
     * @param id 号码内容
     * @return 是否有效
     */
    public static boolean isValid(String id){
        Boolean validResult = true;
        //校验长度只能为15或18
        int len = id.length();
        if (len != FIRST_ID_CARD_LENGTH && len != SECOND_ID_CARD_LENGTH){
            validResult = false;
        }
        //校验生日
        if (!validDate(id)){
            validResult = false;
        }
        return validResult;
    }

    /**
     * 校验生日
     * @param id
     * @return
     */
    private static boolean validDate(String id)
    {
        try
        {
            String birth = id.length() == 15 ? "19" + id.substring(6, 12) : id.substring(6, 14);
            SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMdd");
            Date birthDate = sdf.parse(birth);
            if (!birth.equals(sdf.format(birthDate))){
                return false;
            }
        }
        catch (ParseException e)
        {
            return false;
        }
        return true;
    }



}
