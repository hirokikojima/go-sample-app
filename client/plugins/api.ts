import { axios } from './axios'

export class Api {
  private axios: any;

  constructor(axios: any) {
    this.axios = axios
  }

  /**
   * ユーザ情報を取得する
   * @param id 
   */
  public getUser(id: number) {
    return this.axios.$get('/user/' + id);
  }

  /**
   * サービス一覧を取得する
   * @param userId 
   */
  public getServices() {
    return this.axios.$get('/service')
  }

  /**
   * ユーザのサービス一覧を取得する
   * @param userId 
   */
  public getServicesByUserId(userId: number) {
    return this.axios.$get('/service/' + userId)
  }
}

export default function ({ $axios }: { $axios: any }, inject: any) {
  inject('api', new Api($axios))
}