import styles from './index.less';
import React, { useEffect, useState } from 'react';
import { Button, Space, Input, message, Radio } from 'antd';
import QRCode from 'qrcode';
import copy from 'copy-to-clipboard';
import { CopyOutlined } from '@ant-design/icons';
import { createUrl } from '@/api/url';

export default function IndexPage() {
  const [isLoading, setIsLoading] = useState(false);
  const [imgSrc, setImgSrc] = useState('');
  const [inputValue, setInputValue] = useState<any>();
  const [originUrl, setOriginUrl] = useState('待输入');
  const [shortUrl, setShortUrl] = useState('');
  const [radioDomain, setRadioDomain] = useState('s.shetuankaoqin.com');

  const onSearch = (value: any) => {

    console.log(value);
    let reg = /^((https|http|ftp|rtsp|mms)?:\/\/)[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/;
    if (!value) {
      message.info('请输入待缩短链接地址');
      return;
    } else if (!reg.test(value)) {
      message.error('请输入合法网址，http:// 或 https:// 开头');
      return;
    }

    setOriginUrl(value);
    setIsLoading(true);
    createUrl({
      url: value,
      short_url_prefix: radioDomain
    }).then((res: any) => {
      console.log(res);
      const resUrl = res?.data?.short_url;
      if (resUrl) {
        const safeDomain = `https://${resUrl}`;
        setShortUrl(safeDomain);
        QRCode.toDataURL(safeDomain)
          .then((url: any) => {
            setImgSrc(url);
            message.success('成功生成短链')
          })
          .catch((err: any) => {
            console.error(err)
          })
      } else {
        message.error('生成短链失败');
      }
    }).finally(() => {
      setIsLoading(false);
      setInputValue(undefined);
    }).catch(e => {
      message.error('生成短链失败' + JSON.stringify(e));
    })
  }

  // useEffect(() => {
  // }, []);

  const handleCopy = () => {
    copy(shortUrl);
    message.success('成功复制到粘贴板');
  }

  return (
    <div>
      <div className={styles.headerContainer}>
        <div className={styles.headerContent}>
          <a href="/" title="短链接">
            <img style={{ height: 50 }} src="https://luckteam.oss-cn-beijing.aliyuncs.com/static/img/logo-design.png" alt="短链接" />
          </a>
          {/* <Space className={styles.right}>
            <Button size={'middle'}>营销 Demo</Button>
            <Button type="primary" size={'middle'}>免费使用</Button>
          </Space> */}
        </div>
      </div>
      <div className={styles.mainContainer}>
        <h2 className={styles.slogan}>简单易用的营销短链接</h2>
        <Input.Search
          size='large'
          placeholder="请输入 http:// 或 https:// 开头的网址"
          allowClear
          value={inputValue}
          onChange={e => setInputValue(e.target.value)}
          onSearch={onSearch}
          style={{ width: 600 }}
          loading={isLoading}
          readOnly={isLoading}
          enterButton={'生成短链接'}
        />

        <div className={styles.domainContainer}>
          <p>短链域名：</p>
          <Radio.Group onChange={e => setRadioDomain(e.target.value)} value={radioDomain}>
            <Radio value={'s.shetuankaoqin.com'}>s.shetuankaoqin.com</Radio>
            <Radio value={'x-url.cc'}>x-url.cc</Radio>
            <Radio value={'luckurl.cn'}>luckurl.cn</Radio>
          </Radio.Group>
        </div>

        <div className={styles.urlResult}>
          <b>短链生成结果：</b>
          <div className={styles.resultContent}>
            {imgSrc ? <img src={imgSrc} alt="短链二维码" className={styles.qrCodeImg} />
              : <svg className={styles.defaultImg} viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2711" width="200" height="200"><path d="M85.312 85.312V384H384V85.312H85.312zM0 0h469.248v469.248H0V0z m170.624 170.624h128v128h-128v-128zM0 554.624h469.248v469.248H0V554.624z m85.312 85.312v298.624H384V639.936H85.312z m85.312 85.312h128v128h-128v-128zM554.624 0h469.248v469.248H554.624V0z m85.312 85.312V384h298.624V85.312H639.936z m383.936 682.56H1024v85.376h-298.752V639.936H639.936V1023.872H554.624V554.624h255.936v213.248h128V554.624h85.312v213.248z m-298.624-597.248h128v128h-128v-128z m298.624 853.248h-85.312v-85.312h85.312v85.312z m-213.312 0h-85.312v-85.312h85.312v85.312z" fill="#262626" p-id="2712"></path></svg>}
            <div className={styles.urlContent}>
              <div className={styles.shortUrl}>短链接：
                {
                  shortUrl ?
                    <Space>
                      <a href={shortUrl} target='_blank'>{shortUrl}</a>
                      <Button type="primary" shape="circle" icon={<CopyOutlined />} onClick={handleCopy} />
                    </Space>
                    : <span className={styles.originUrl}>待生成</span>
                }
              </div>
              <p title={originUrl} className={styles.originUrl}>原始链接：{originUrl}</p>
            </div>
          </div>
          <p>注意：此短链接 <span style={{ color: '#3464e0' }}>永久有效</span>
          {/* ，登录可查看完整访问数据 */}
          </p>
        </div>
        {/* <img src='https://luckteam.oss-cn-beijing.aliyuncs.com/static/img/bg.png' /> */}
      </div>
    </div>
  );
}
